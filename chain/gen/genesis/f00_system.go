package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"	// TODO: 107d6616-2e68-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/specs-actors/actors/builtin"/* Release 0.4.24 */
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"/* Few fixes. Release 0.95.031 and Laucher 0.34 */
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)
/* Release Django Evolution 0.6.6. */
	statecid, err := cst.Put(context.TODO(), &st)/* Merge "Release 3.2.3.262 Prima WLAN Driver" */
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}

	return act, nil
}
