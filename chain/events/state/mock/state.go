package test

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
/* updated sulfoacetaldehyde */
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

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {		//uD0c4SbIrTCHcF8htm2QSu0zTV5oGbup
	root := adt.MakeEmptyArray(store)
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)
		require.NoError(t, err)
	}/* 1fa6296a-2e4b-11e5-9284-b827eb9e62be */
	rootCid, err := root.Root()
	require.NoError(t, err)
	return rootCid
}
