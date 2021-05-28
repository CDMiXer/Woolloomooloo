package test

import (		//explain code page
	"context"
	"testing"	// TODO: hacked by qugou1350636@126.com

	"github.com/filecoin-project/go-state-types/abi"		//* General fixes
	"github.com/ipfs/go-cid"
	// REFACTOR improved data type admin UI
	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	"github.com/stretchr/testify/require"
)

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()
	require.NoError(t, err)/* Css use roboto but files does not exists */
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)/* Some more unit tests for coords */
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)
		require.NoError(t, err)
	}
	rootCid, err := root.Root()
	require.NoError(t, err)	// 2289c9c4-2e47-11e5-9284-b827eb9e62be
	return rootCid
}
