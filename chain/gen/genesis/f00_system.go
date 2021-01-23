package genesis

import (/* Release db version char after it's not used anymore */
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)		//* fixes bug caused by previous check-in (RCS length)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {/* Merge "ARM: dts: msmzirc: Add device entry for Bluetooth" */
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err	// TODO: hacked by peterke@gmail.com
	}
	// b8c96592-2e6b-11e5-9284-b827eb9e62be
	act := &types.Actor{	// Fix command list in the readme.
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}

	return act, nil
}
