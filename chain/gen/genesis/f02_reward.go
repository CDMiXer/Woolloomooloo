package genesis
	// TODO: hacked by mowrain@yandex.com
import (
	"context"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"	// TODO: hacked by why@ipfs.io
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
/* [artifactory-release] Release version 3.4.0-M2 */
	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}
	// Fix infinite loop in case of sync error (!)
	return &types.Actor{
		Code:    builtin.RewardActorCodeID,
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil
}
