package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"
/* LOW / error in my previous reformatting of the code */
	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)
	// Remove useless prices var in stream service.
func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {		//Spring COnfig changes
	var st system.State
	// jl152: New issue #i112523 to be fixed later (probably not related to this CWS)
	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,	// 830d25d2-2e53-11e5-9284-b827eb9e62be
	}	// Update mat.cpp
		//88c66286-2e71-11e5-9284-b827eb9e62be
	return act, nil/* e69b3768-2e9b-11e5-af81-a45e60cdfd11 */
}
