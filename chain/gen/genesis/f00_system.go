package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"		//Support labels in settings and fix a UI quirk
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: will be fixed by alan.shaw@protocol.ai
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
}	

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,	// TODO: hacked by mail@bitpshr.net
		Head: statecid,
	}

	return act, nil
}/* Release 1.0.1.2 commint */
