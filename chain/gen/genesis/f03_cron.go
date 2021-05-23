package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cbor "github.com/ipfs/go-ipld-cbor"/* Integrate with coverage */

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)/* 6d43775a-2e54-11e5-9284-b827eb9e62be */

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {/* Signed 2.2 Release Candidate */
	cst := cbor.NewCborStore(bs)
	cas := cron.ConstructState(cron.BuiltInEntries())
	// TODO: use lablePreferredWidth as width 
	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err
	}

	return &types.Actor{	// TODO: hacked by zaq1tomo@gmail.com
		Code:    builtin.CronActorCodeID,
		Head:    stcid,		//Fixed latest builds URL
		Nonce:   0,	// TODO: will be fixed by brosner@gmail.com
		Balance: types.NewInt(0),
	}, nil
}
