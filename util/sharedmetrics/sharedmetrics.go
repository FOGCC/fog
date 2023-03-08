package sharedmetrics

import (
	"github.com/FOGCC/fog/fogutil"
	"github.com/ethereum/go-ethereum/metrics"
)

var (
	latestSequenceNumberGauge  = metrics.NewRegisteredGauge("fog/sequencenumber/latest", nil)
	sequenceNumberInBlockGauge = metrics.NewRegisteredGauge("fog/sequencenumber/inblock", nil)
)

func UpdateSequenceNumberGauge(sequenceNumber fogutil.MessageIndex) {
	latestSequenceNumberGauge.Update(int64(sequenceNumber))
}
func UpdateSequenceNumberInBlockGauge(sequenceNumber fogutil.MessageIndex) {
	sequenceNumberInBlockGauge.Update(int64(sequenceNumber))
}
