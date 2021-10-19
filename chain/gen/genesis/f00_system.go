package genesis
/* (tanner) Release 1.14rc2 */
import (
	"context"
	// disable nginx access logs for now
	"github.com/filecoin-project/specs-actors/actors/builtin/system"
	// Update StepImplementation.cs
	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"		//Forgot vector doesn't automatically resize when just using operator[]

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State
/* add IBM Swift Sandbox (REPL) to iOS section */
	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}
	// Creating directory for 2.8 goodness.
	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}

	return act, nil
}/* Update UI and remove RSS feed. */
