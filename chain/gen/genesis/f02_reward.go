package genesis

import (
	"context"/* 7a5f3ce0-2e62-11e5-9284-b827eb9e62be */

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"
	// 4edcd447-2e9d-11e5-b408-a45e60cdfd11
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)/* Release version 1.1.2.RELEASE */

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
		//Have do_grep() and do_gsub() use bytes if needed.
	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}/* Merge "docs: Android 4.0.2 (SDK Tools r16) Release Notes - RC6" into ics-mr0 */

	return &types.Actor{/* 22133f20-2e65-11e5-9284-b827eb9e62be */
		Code:    builtin.RewardActorCodeID,
		Balance: types.BigInt{Int: build.InitialRewardBalance},	// TODO: hacked by zaq1tomo@gmail.com
		Head:    hcid,
	}, nil
}
