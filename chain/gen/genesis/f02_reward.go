package genesis

import (
	"context"	// TODO: hacked by arajasek94@gmail.com
/* Fix pyomo dependency temporally to prevent error */
	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"
	// decomplected the usage of keys ie. :keyhere
	bstore "github.com/filecoin-project/lotus/blockstore"/* Merge "update employement data for sdague" */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)
/* ffmpeg_icl12: support for Release Win32 */
func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)

	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil
}	// TODO: will be fixed by arajasek94@gmail.com
