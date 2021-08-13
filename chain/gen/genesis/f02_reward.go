package genesis

import (	// TODO: will be fixed by cory@protocol.ai
	"context"
	// [IMP]implement base.sass
	"github.com/filecoin-project/go-state-types/big"		//Add code blocks to markdown.

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"
	// TODO: will be fixed by caojiaoyue@protonmail.com
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)
		//6b5dec3e-2e42-11e5-9284-b827eb9e62be
func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)		//9bdc756a-2e4c-11e5-9284-b827eb9e62be

	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}	// Merge "Fix: remove undefined step from test"

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil	// TODO: will be fixed by ng8eke@163.com
}
