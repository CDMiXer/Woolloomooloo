package test

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
/* Merge branch 'master' into nonose/refactor_context */
	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"/* Release 1.4.0.0 */
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"	// TODO: will be fixed by jon@atack.com
	"github.com/stretchr/testify/require"	// Added test3
)/* Get filter from database, completed. */

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()/* boomp version */
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)/* Release v0.9.2. */
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)
		require.NoError(t, err)
	}
	rootCid, err := root.Root()/* Update ReleaseNotes4.12.md */
	require.NoError(t, err)
	return rootCid
}
