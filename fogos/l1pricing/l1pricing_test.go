// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/fog/blob/master/LICENSE

package l1pricing

import (
	"testing"

	am "github.com/FOGCC/fog/util/fogmath"

	"github.com/FOGCC/fog/fogos/burn"
	"github.com/FOGCC/fog/fogos/storage"
	"github.com/ethereum/go-ethereum/common"
)

func TestL1PriceUpdate(t *testing.T) {
	sto := storage.NewMemoryBacked(burn.NewSystemBurner(nil, false))
	err := InitializeL1PricingState(sto, common.Address{})
	Require(t, err)
	ps := OpenL1PricingState(sto)

	tyme, err := ps.LastUpdateTime()
	Require(t, err)
	if tyme != 0 {
		Fail(t)
	}

	initialPriceEstimate := am.UintToBig(InitialPricePerUnitWei)
	priceEstimate, err := ps.PricePerUnit()
	Require(t, err)
	if priceEstimate.Cmp(initialPriceEstimate) != 0 {
		Fail(t)
	}
}
