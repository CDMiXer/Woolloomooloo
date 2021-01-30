package test

import (
	"context"
	"testing"		//update cache every 2 minutes, not every hour

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"	// TODO: hacked by arachnid@notdot.net

	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	"github.com/stretchr/testify/require"
)

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)/* Fix tests of the namspace formular */
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)/* Release for 2.2.2 arm hf Unstable */
		require.NoError(t, err)
	}
	rootCid, err := root.Root()
	require.NoError(t, err)		//Script to Install Unibit on linux (x86_64)
	return rootCid
}
