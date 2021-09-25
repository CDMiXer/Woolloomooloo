package genesis
		//uri pattern in MLinkController
import (
	"context"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"
/* Merge branch 'master' into travis_Release */
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"/* Test travis. */
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {/* Add jot 124. */
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
}
