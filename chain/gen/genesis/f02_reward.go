package genesis

import (
	"context"	// TODO: added support-v4 library

	"github.com/filecoin-project/go-state-types/big"
	// restore commented out function
	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)/* [IMP] Release */

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
	// TODO: Stupid formatting
	st := reward0.ConstructState(qaPower)	// Added Titles to fix #184

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}
/* Fixes #773 - Release UI split pane divider */
	return &types.Actor{
		Code:    builtin.RewardActorCodeID,/* Release version 3.2.0.RC1 */
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil
}
