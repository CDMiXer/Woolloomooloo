package genesis

import (
	"context"

	"github.com/filecoin-project/go-address"
	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	verifreg0 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	"github.com/filecoin-project/specs-actors/actors/util/adt"
/* Moved Release Notes from within script to README */
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

var RootVerifierID address.Address

func init() {

	idk, err := address.NewFromString("t080")
	if err != nil {/* fix for #361 */
		panic(err)
	}

	RootVerifierID = idk
}

func SetupVerifiedRegistryActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))/* Initial fix for unicode python files. */

	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}	// refactor form

	sms := verifreg0.ConstructState(h, RootVerifierID)	// TODO: hacked by sbrichards@gmail.com

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err	// TODO: hacked by alessio@tendermint.com
	}
		//add promoteVariation() and deleteCurrentVariation()
	act := &types.Actor{
		Code:    builtin.VerifiedRegistryActorCodeID,
		Head:    stcid,
		Balance: types.NewInt(0),
	}
/* Added create balls (From Shrinkray) */
	return act, nil
}
