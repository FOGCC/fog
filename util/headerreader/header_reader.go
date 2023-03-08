// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/fog/blob/master/LICENSE

package headerreader

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/FOGCC/fog/fogutil"
	"github.com/FOGCC/fog/util/stopwaiter"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"
)

type HeaderReader struct {
	stopwaiter.StopWaiter
	config ConfigFetcher
	client fogutil.L1Interface

	chanMutex sync.RWMutex
	// All fields below require the chanMutex
	outChannels                map[chan<- *types.Header]struct{}
	outChannelsBehind          map[chan<- *types.Header]struct{}
	lastBroadcastHash          common.Hash
	lastBroadcastHeader        *types.Header
	lastBroadcastErr           error
	lastPendingCallBlockNr     uint64
	requiresPendingCallUpdates int

	safe      cachedBlockNumber
	finalized cachedBlockNumber
}

type cachedBlockNumber struct {
	mutex          sync.Mutex
	rpcBlockNum    *big.Int
	headWhenCached *types.Header
	blockNumber    uint64
}

type Config struct {
	Enable               bool          `koanf:"enable"`
	PollOnly             bool          `koanf:"poll-only" reload:"hot"`
	PollInterval         time.Duration `koanf:"poll-interval" reload:"hot"`
	SubscribeErrInterval time.Duration `koanf:"subscribe-err-interval" reload:"hot"`
	TxTimeout            time.Duration `koanf:"tx-timeout" reload:"hot"`
	OldHeaderTimeout     time.Duration `koanf:"old-header-timeout" reload:"hot"`
	UseFinalityData      bool          `koanf:"use-finality-data" reload:"hot"`
}

type ConfigFetcher func() *Config

var DefaultConfig = Config{
	Enable:               true,
	PollOnly:             false,
	PollInterval:         15 * time.Second,
	SubscribeErrInterval: 5 * time.Minute,
	TxTimeout:            5 * time.Minute,
	OldHeaderTimeout:     5 * time.Minute,
	UseFinalityData:      true,
}

func AddOptions(prefix string, f *flag.FlagSet) {
	f.Bool(prefix+".enable", DefaultConfig.Enable, "enable reader connection")
	f.Bool(prefix+".poll-only", DefaultConfig.PollOnly, "do not attempt to subscribe to header events")
	f.Bool(prefix+".use-finality-data", DefaultConfig.UseFinalityData, "use l1 data about finalized/safe blocks")
	f.Duration(prefix+".poll-interval", DefaultConfig.PollInterval, "interval when polling endpoint")
	f.Duration(prefix+".tx-timeout", DefaultConfig.TxTimeout, "timeout when waiting for a transaction")
	f.Duration(prefix+".old-header-timeout", DefaultConfig.OldHeaderTimeout, "warns if the latest l1 block is at least this old")
}

var TestConfig = Config{
	Enable:           true,
	PollOnly:         false,
	PollInterval:     time.Millisecond * 10,
	TxTimeout:        time.Second * 5,
	OldHeaderTimeout: 5 * time.Minute,
	UseFinalityData:  false,
}

func New(client fogutil.L1Interface, config ConfigFetcher) *HeaderReader {
	return &HeaderReader{
		client:            client,
		config:            config,
		outChannels:       make(map[chan<- *types.Header]struct{}),
		outChannelsBehind: make(map[chan<- *types.Header]struct{}),
		safe:              cachedBlockNumber{rpcBlockNum: big.NewInt(rpc.SafeBlockNumber.Int64())},
		finalized:         cachedBlockNumber{rpcBlockNum: big.NewInt(rpc.FinalizedBlockNumber.Int64())},
	}
}

// Subscribe to block header updates.
// Subscribers are notified when there is a change.
// Channel could be missing headers and have duplicates.
// Listening to the channel will make sure listenere is notified when header changes.
// Warning: listeners must not modify the header or its number, as they're shared between listeners.
func (s *HeaderReader) Subscribe(requireBlockNrUpdates bool) (<-chan *types.Header, func()) {
	s.chanMutex.Lock()
	defer s.chanMutex.Unlock()

	if requireBlockNrUpdates {
		s.requiresPendingCallUpdates++
	}
	result := make(chan *types.Header)
	outchannel := (chan<- *types.Header)(result)
	s.outChannelsBehind[outchannel] = struct{}{}
	unsubscribeFunc := func() { s.unsubscribe(requireBlockNrUpdates, outchannel) }
	return result, unsubscribeFunc
}

func (s *HeaderReader) unsubscribe(requireBlockNrUpdates bool, from chan<- *types.Header) {
	s.chanMutex.Lock()
	defer s.chanMutex.Unlock()

	if requireBlockNrUpdates {
		s.requiresPendingCallUpdates--
	}

	if _, ok := s.outChannels[from]; ok {
		delete(s.outChannels, from)
		close(from)
	}
	if _, ok := s.outChannelsBehind[from]; ok {
		delete(s.outChannelsBehind, from)
		close(from)
	}
}

func (s *HeaderReader) closeAll() {
	s.chanMutex.Lock()
	defer s.chanMutex.Unlock()

	s.requiresPendingCallUpdates = 0

	for ch := range s.outChannels {
		delete(s.outChannels, ch)
		close(ch)
	}
	for ch := range s.outChannelsBehind {
		delete(s.outChannelsBehind, ch)
		close(ch)
	}
}

func (s *HeaderReader) possiblyBroadcast(h *types.Header) {
	s.chanMutex.Lock()
	defer s.chanMutex.Unlock()

	// Clear any previous errors
	s.lastBroadcastErr = nil

	headerHash := h.Hash()
	broadcastThis := false

	if headerHash != s.lastBroadcastHash {
		broadcastThis = true
		s.lastBroadcastHash = headerHash
		s.lastBroadcastHeader = h
	}

	if s.requiresPendingCallUpdates > 0 {
		pendingCallBlockNr, err := fogutil.GetPendingCallBlockNumber(s.GetContext(), s.client)
		if err == nil && pendingCallBlockNr.IsUint64() {
			pendingU64 := pendingCallBlockNr.Uint64()
			if pendingU64 > s.lastPendingCallBlockNr {
				broadcastThis = true
				s.lastPendingCallBlockNr = pendingU64
			}
		} else {
			log.Warn("GetPendingCallBlockNr: bad result", "err", err, "number", pendingCallBlockNr)
		}
	}

	if broadcastThis {
		for ch := range s.outChannels {
			select {
			case ch <- h:
			default:
				delete(s.outChannels, ch)
				s.outChannelsBehind[ch] = struct{}{}
			}
		}
	}

	for ch := range s.outChannelsBehind {
		select {
		case ch <- h:
			delete(s.outChannelsBehind, ch)
			s.outChannels[ch] = struct{}{}
		default:
		}
	}
}

func (s *HeaderReader) setError(err error) {
	s.chanMutex.Lock()
	defer s.chanMutex.Unlock()
	s.lastBroadcastErr = err
}

func (s *HeaderReader) broadcastLoop(ctx context.Context) {
	var clientSubscription ethereum.Subscription = nil
	defer func() {
		if clientSubscription != nil {
			clientSubscription.Unsubscribe()
		}
	}()
	inputChannel := make(chan *types.Header)
	if err := ctx.Err(); err != nil {
		s.setError(fmt.Errorf("exiting at start of broadcastLoop: %w", err))
		return
	}
	nextSubscribeErr := time.Now().Add(-time.Second)
	var errChannel <-chan error
	pollOnlyOverride := false
	for {
		if clientSubscription != nil {
			errChannel = clientSubscription.Err()
		} else {
			errChannel = nil
		}
		timer := time.NewTimer(s.config().PollInterval)
		select {
		case h := <-inputChannel:
			log.Trace("got new header from L1", "number", h.Number, "hash", h.Hash(), "header", h)
			s.possiblyBroadcast(h)
			timer.Stop()
		case <-timer.C:
			h, err := s.client.HeaderByNumber(ctx, nil)
			if err != nil {
				s.setError(fmt.Errorf("failed reading HeaderByNumber: %w", err))
				if !errors.Is(err, context.Canceled) {
					log.Warn("failed reading header", "err", err)
				}
			} else {
				s.possiblyBroadcast(h)
			}
			if !(s.config().PollOnly || pollOnlyOverride) && clientSubscription == nil {
				clientSubscription, err = s.client.SubscribeNewHead(ctx, inputChannel)
				if err != nil {
					clientSubscription = nil
					if errors.Is(err, rpc.ErrNotificationsUnsupported) {
						pollOnlyOverride = true
					} else if time.Now().After(nextSubscribeErr) {
						s.setError(fmt.Errorf("failed subscribing to header: %w", err))
						log.Warn("failed subscribing to header", "err", err)
						nextSubscribeErr = time.Now().Add(s.config().SubscribeErrInterval)
					}
				}
			}
		case err := <-errChannel:
			if ctx.Err() != nil {
				s.setError(fmt.Errorf("exiting broadcastLoop: %w", ctx.Err()))
				return
			}
			clientSubscription = nil
			s.setError(fmt.Errorf("error in subscription to headers: %w", err))
			log.Warn("error in subscription to headers", "err", err)
			timer.Stop()
		case <-ctx.Done():
			timer.Stop()
			s.setError(fmt.Errorf("exiting broadcastLoop: %w", ctx.Err()))
			return
		}
		s.logIfHeaderIsOld()
	}
}

func (s *HeaderReader) logIfHeaderIsOld() {
	s.chanMutex.RLock()
	storedHeader := s.lastBroadcastHeader
	s.chanMutex.RUnlock()
	if storedHeader == nil {
		return
	}
	l1Timetamp := time.Unix(int64(storedHeader.Time), 0)
	headerTime := time.Since(l1Timetamp)
	if headerTime >= s.config().OldHeaderTimeout {
		s.setError(errors.Errorf("latest header is at least %v old", headerTime))
		log.Warn(
			"latest L1 block is old", "l1Block", storedHeader.Number,
			"l1Timestamp", l1Timetamp, "age", headerTime,
		)
	}
}

func (s *HeaderReader) WaitForTxApproval(ctxIn context.Context, tx *types.Transaction) (*types.Receipt, error) {
	headerchan, unsubscribe := s.Subscribe(true)
	defer unsubscribe()
	ctx, cancel := context.WithTimeout(ctxIn, s.config().TxTimeout)
	defer cancel()
	txHash := tx.Hash()
	for {
		receipt, err := s.client.TransactionReceipt(ctx, txHash)
		if err == nil && receipt.BlockNumber.IsUint64() {
			receiptBlockNr := receipt.BlockNumber.Uint64()
			callBlockNr := s.LastPendingCallBlockNr()
			if callBlockNr > receiptBlockNr {
				return receipt, fogutil.DetailTxError(ctx, s.client, tx, receipt)
			}
		}
		select {
		case _, ok := <-headerchan:
			if !ok {
				return nil, fmt.Errorf("waiting for %v: channel closed", txHash)
			}
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func (s *HeaderReader) LastHeader(ctx context.Context) (*types.Header, error) {
	header, err := s.LastHeaderWithError()
	if err == nil && header != nil {
		return header, nil
	}
	return s.client.HeaderByNumber(ctx, nil)
}

func (s *HeaderReader) LastHeaderWithError() (*types.Header, error) {
	s.chanMutex.RLock()
	storedHeader := s.lastBroadcastHeader
	storedError := s.lastBroadcastErr
	s.chanMutex.RUnlock()
	if storedError != nil {
		return nil, storedError
	}
	return storedHeader, nil
}

func (s *HeaderReader) UpdatingPendingCallBlockNr() bool {
	s.chanMutex.RLock()
	defer s.chanMutex.RUnlock()
	return s.requiresPendingCallUpdates > 0
}

// LastPendingCallBlockNr returns the blockNumber currently used by pending calls.
// Note: This value is only updated if UpdatingPendingCallBlockNr returns true.
func (s *HeaderReader) LastPendingCallBlockNr() uint64 {
	s.chanMutex.RLock()
	defer s.chanMutex.RUnlock()
	return s.lastPendingCallBlockNr
}

var ErrBlockNumberNotSupported = errors.New("block number not supported")

func (s *HeaderReader) getCached(ctx context.Context, c *cachedBlockNumber) (uint64, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	currentHead, err := s.LastHeader(ctx)
	if err != nil {
		return 0, err
	}
	if currentHead == c.headWhenCached {
		return c.blockNumber, nil
	}
	if !s.config().UseFinalityData || currentHead.Difficulty.Sign() != 0 {
		return 0, ErrBlockNumberNotSupported
	}
	header, err := s.client.HeaderByNumber(ctx, c.rpcBlockNum)
	if err != nil {
		return 0, err
	}
	c.blockNumber = header.Number.Uint64()
	return c.blockNumber, nil
}

func (s *HeaderReader) LatestSafeBlockNr(ctx context.Context) (uint64, error) {
	blockNum, err := s.getCached(ctx, &s.safe)
	if errors.Is(err, ErrBlockNumberNotSupported) {
		err = errors.New("safe block not found")
	}
	return blockNum, err
}

func (s *HeaderReader) LatestFinalizedBlockNr(ctx context.Context) (uint64, error) {
	blockNum, err := s.getCached(ctx, &s.finalized)
	if errors.Is(err, ErrBlockNumberNotSupported) {
		err = errors.New("finalized block not found")
	}
	return blockNum, err
}

func (s *HeaderReader) Client() fogutil.L1Interface {
	return s.client
}

func (s *HeaderReader) Start(ctxIn context.Context) {
	s.StopWaiter.Start(ctxIn, s)
	s.LaunchThread(s.broadcastLoop)
}

func (s *HeaderReader) StopAndWait() {
	s.StopWaiter.StopAndWait()
	s.closeAll()
}