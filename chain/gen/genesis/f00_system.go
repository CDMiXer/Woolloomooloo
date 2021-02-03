package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"
/* Changed host list layout; refactoring */
	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err/* Add simple watching to documentation */
	}
/* Merge branch 'master' into add_Mohamed_Gomaa */
	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}

	return act, nil
}	// TODO: hacked by remco@dutchcoders.io
