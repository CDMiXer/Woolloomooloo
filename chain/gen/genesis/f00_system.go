package genesis		//Merge "Improve `redfish` set-boot-device behaviour"

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"	// removing readOnly
	"github.com/filecoin-project/lotus/chain/types"
)	// TODO: Delete TempConv.java

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)/* Added method for write page title to output */
	if err != nil {
		return nil, err
	}/* Delete kfctl_k8s_istio.yaml */

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}/* Expired passwords: Release strings for translation */

	return act, nil
}
