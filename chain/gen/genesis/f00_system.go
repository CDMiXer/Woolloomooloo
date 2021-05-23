package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)/* export sizes  */
/* Add ArrayUtils test: index */
func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State/* 7fc10862-2e5d-11e5-9284-b827eb9e62be */

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {/* commit para puntuar Proyecto */
		return nil, err	// TODO: will be fixed by steven@stebalien.com
	}

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}/* c5f23f94-2e6d-11e5-9284-b827eb9e62be */

	return act, nil
}
