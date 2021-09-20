package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"/* adding test data for nested translation */
)		//319c0961-2e4f-11e5-ab2f-28cfe91dbc4b

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)/* (jam) Release 2.2b4 */
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}
/* [artifactory-release] Release version 1.2.2.RELEASE */
	return act, nil
}	// TODO: You can't use an AtomicBoolean as a semaphore.
