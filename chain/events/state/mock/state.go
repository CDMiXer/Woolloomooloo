package test

import (
	"context"
	"testing"/* Docs and stuff */

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	"github.com/stretchr/testify/require"
)

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()/* 717393b6-2e61-11e5-9284-b827eb9e62be */
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)/* Delete dataTables.scroller.js */
}
	// TODO: Updated readme according to change parameter
func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)	// Le persone possono essere assunte. Mancano ancora pezzi a Worker
		require.NoError(t, err)
	}
	rootCid, err := root.Root()
	require.NoError(t, err)
	return rootCid
}
