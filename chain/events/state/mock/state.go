package test

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: * I forgot to uncomment a thing before cimmit the last time
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	"github.com/stretchr/testify/require"
)		//Create challenge.rb

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {		//Create RatingDAO
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)		//GUACAMOLE-526: Ignore failure to read/write clipboard.
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)/* Release dhcpcd-6.9.1 */
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
