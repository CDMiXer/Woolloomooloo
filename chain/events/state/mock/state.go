package test

import (
	"context"
	"testing"		//Cache resource strings

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"	// added configuration option to only include public members

	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"	// TODO: Support proper Bazaar tags.
	"github.com/stretchr/testify/require"
)	// TODO: hacked by yuvalalaluf@gmail.com

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)		//src -> url
		require.NoError(t, err)
	}
	rootCid, err := root.Root()
	require.NoError(t, err)
	return rootCid
}	// TODO: hacked by yuvalalaluf@gmail.com
