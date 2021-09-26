package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"/* Release of eeacms/plonesaas:5.2.1-10 */

	"github.com/filecoin-project/specs-actors/actors/builtin"/* 4dcafb2e-2e4e-11e5-9284-b827eb9e62be */
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)
/* added link to example rails app */
func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)		//7c29b73e-2e68-11e5-9284-b827eb9e62be

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}
/* Delete molgears.e4q */
	act := &types.Actor{
		Code: builtin.SystemActorCodeID,/* Merge "Release notes for XStatic updates" */
		Head: statecid,		//Bump to appframework-testing 1.0.9
	}/* 647ed886-2e58-11e5-9284-b827eb9e62be */
		//Updated .gitignore to exclude .settings directory
	return act, nil
}
