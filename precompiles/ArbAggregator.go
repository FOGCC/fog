// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/fog/blob/master/LICENSE

package precompiles

import (
	"errors"
	"math/big"

	"github.com/FOGCC/fog/fogos/l1pricing"
)

// fogAggregator provides aggregators and their users methods for configuring how they participate in L1 aggregation.
// FOG One's default aggregator is the Sequencer, which a user will prefer unless SetPreferredAggregator()
// is invoked to change it.
type fogAggregator struct {
	Address addr // 0x6d
}

var ErrNotOwner = errors.New("must be called by chain owner")

// GetPreferredAggregator returns the preferred aggregator address.
// Deprecated: Do not use this method.
func (con fogAggregator) GetPreferredAggregator(c ctx, evm mech, address addr) (prefAgg addr, isDefault bool, err error) {
	return l1pricing.BatchPosterAddress, true, err
}

// GetDefaultAggregator returns the default aggregator address.
// Deprecated: Do not use this method.
func (con fogAggregator) GetDefaultAggregator(c ctx, evm mech) (addr, error) {
	return l1pricing.BatchPosterAddress, nil
}

// GetBatchPosters gets the addresses of all current batch posters
func (con fogAggregator) GetBatchPosters(c ctx, evm mech) ([]addr, error) {
	return c.State.L1PricingState().BatchPosterTable().AllPosters(65536)
}

func (con fogAggregator) AddBatchPoster(c ctx, evm mech, newBatchPoster addr) error {
	isOwner, err := c.State.ChainOwners().IsMember(c.caller)
	if err != nil {
		return err
	}
	if !isOwner {
		return ErrNotOwner
	}
	batchPosterTable := c.State.L1PricingState().BatchPosterTable()
	isBatchPoster, err := batchPosterTable.ContainsPoster(newBatchPoster)
	if err != nil {
		return err
	}
	if !isBatchPoster {
		_, err = batchPosterTable.AddPoster(newBatchPoster, newBatchPoster)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetFeeCollector gets a batch poster's fee collector
func (con fogAggregator) GetFeeCollector(c ctx, evm mech, batchPoster addr) (addr, error) {
	posterInfo, err := c.State.L1PricingState().BatchPosterTable().OpenPoster(batchPoster, false)
	if err != nil {
		return addr{}, err
	}
	return posterInfo.PayTo()
}

// SetFeeCollector sets a batch poster's fee collector (caller must be the batch poster, its fee collector, or an owner)
func (con fogAggregator) SetFeeCollector(c ctx, evm mech, batchPoster addr, newFeeCollector addr) error {
	posterInfo, err := c.State.L1PricingState().BatchPosterTable().OpenPoster(batchPoster, false)
	if err != nil {
		return err
	}
	oldFeeCollector, err := posterInfo.PayTo()
	if err != nil {
		return err
	}
	if c.caller != batchPoster && c.caller != oldFeeCollector {
		isOwner, err := c.State.ChainOwners().IsMember(c.caller)
		if err != nil {
			return err
		}
		if !isOwner {
			return errors.New("only a batch poster (or its fee collector / chain owner) may change its fee collector")
		}
	}
	return posterInfo.SetPayTo(newFeeCollector)
}

// GetTxBaseFee gets an aggregator's current fixed fee to submit a tx
func (con fogAggregator) GetTxBaseFee(c ctx, evm mech, aggregator addr) (huge, error) {
	// This is deprecated and now always returns zero.
	return big.NewInt(0), nil
}

// SetTxBaseFee sets an aggregator's fixed fee (caller must be the aggregator, its fee collector, or an owner)
func (con fogAggregator) SetTxBaseFee(c ctx, evm mech, aggregator addr, feeInL1Gas huge) error {
	// This is deprecated and is now a no-op.
	return nil
}
