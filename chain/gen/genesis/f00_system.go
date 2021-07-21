package genesis

import (/* Release 1.2.8 */
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)		//Updated Dev properties to support SSO and added debuging for configs

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)
	// TODO: update test page with embed parameters
	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
}	

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}	// TODO: hacked by onhardev@bk.ru

	return act, nil
}
