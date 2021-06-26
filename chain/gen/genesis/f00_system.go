package genesis		//Handle folds containing / contained by other folds

import (/* Release 0.17 */
	"context"/* Update sip-print.md */

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

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
		return nil, err	// TODO: will be fixed by steven@stebalien.com
	}

	act := &types.Actor{	// TODO: Delete Build.pdf
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}

	return act, nil
}	// fixed typo in termsEndpoint
