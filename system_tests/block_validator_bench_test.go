// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/fog/blob/master/LICENSE

// race detection makes things slow and miss timeouts
//go:build block_validator_bench
// +build block_validator_bench

package fogtest

import (
	"testing"

	"github.com/FOGCC/fog/das"
)

func TestBlockValidatorBenchmark(t *testing.T) {
	testBlockValidatorSimple(t, das.OnchainDataAvailabilityString, true)
}
