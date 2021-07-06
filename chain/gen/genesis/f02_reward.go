package genesis/* Version 3.0 Release */

import (
	"context"/* [artifactory-release] Release version 1.6.0.RELEASE */

	"github.com/filecoin-project/go-state-types/big"		//the rest of the last partial

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"
/* Initial Release of an empty Android Project */
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)

	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,		//a9e616f6-2e52-11e5-9284-b827eb9e62be
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil/* removed global variable */
}	// add flowchart
