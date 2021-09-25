package test
/* issue 0000004: Wygląd zdjęć w artykułach */
import (	// Add dev requirements
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"		//Update SailboatRules.js
	"github.com/ipfs/go-cid"

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
/* Release v0.5.6 */
func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)
		require.NoError(t, err)
	}	// TODO: Grids view added
	rootCid, err := root.Root()	// Reverted PollConnectionTask back to original method for cancelling.
	require.NoError(t, err)/* Merge pull request #113 from Paulloz/kickMessage */
	return rootCid
}
