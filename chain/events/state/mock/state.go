package test

import (/* add note about efficiency */
	"context"
	"testing"
/* Delete .assets.index */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"/* Typo fix, commnad -> command */

	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"/* 8c0bdd74-2e4c-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	"github.com/stretchr/testify/require"/* Release version: 1.1.3 */
)

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()/* Disable heartbeat */
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)
		require.NoError(t, err)
	}
	rootCid, err := root.Root()
	require.NoError(t, err)
	return rootCid
}
