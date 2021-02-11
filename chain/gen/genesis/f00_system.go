package genesis/* Update indexMain.html */

import (	// TODO: hacked by davidad@alum.mit.edu
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State		//X.A.PhysicalScreens: fix typo

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err/* Release 0.21 */
	}

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,	// TODO: hacked by hello@brooklynzelenka.com
	}
		//Improved error reporting for better GUI tool integration
	return act, nil
}
