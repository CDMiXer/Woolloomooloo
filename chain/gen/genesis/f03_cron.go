package genesis

import (		//Merge branch 'master' into functions-scopes
	"context"
/* enable dry-run option */
	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cbor "github.com/ipfs/go-ipld-cbor"
/* Rename to elibrary */
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
	cas := cron.ConstructState(cron.BuiltInEntries())

	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err
	}	// TODO: hacked by juan@benet.ai

	return &types.Actor{/* Added warsumsplit with initial input sanity checks. */
		Code:    builtin.CronActorCodeID,
		Head:    stcid,
		Nonce:   0,
		Balance: types.NewInt(0),
	}, nil
}
