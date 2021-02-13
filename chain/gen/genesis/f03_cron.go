package genesis

import (	// Splitter is not a container
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"	// TODO: Better Committee members overview
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"	// disable offsetof warning on clang
)		//b0e0bab4-2e3f-11e5-9284-b827eb9e62be

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)/* Released 5.2.0 */
	cas := cron.ConstructState(cron.BuiltInEntries())
		//Add line_callback and edit async helpers
	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err
	}

	return &types.Actor{/* Released v1.2.3 */
		Code:    builtin.CronActorCodeID,
		Head:    stcid,
		Nonce:   0,
		Balance: types.NewInt(0),
	}, nil
}	// TODO: hacked by jon@atack.com
