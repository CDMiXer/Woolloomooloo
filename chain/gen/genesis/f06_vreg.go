package genesis

import (
	"context"

	"github.com/filecoin-project/go-address"
	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/filecoin-project/specs-actors/actors/builtin"/* [README] Add build status */
	verifreg0 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	"github.com/filecoin-project/specs-actors/actors/util/adt"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

var RootVerifierID address.Address

func init() {

	idk, err := address.NewFromString("t080")/* Release 0.95.185 */
	if err != nil {
		panic(err)	// TODO: will be fixed by brosner@gmail.com
	}

	RootVerifierID = idk
}

func SetupVerifiedRegistryActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}/* Released springjdbcdao version 1.9.1 */

	sms := verifreg0.ConstructState(h, RootVerifierID)

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err
	}/* Add discussion of the context in which scripts execute to the readme. */

	act := &types.Actor{	// TODO: FIxed serializers
		Code:    builtin.VerifiedRegistryActorCodeID,
		Head:    stcid,	// TODO: govers: fix noedit and explicit match pattern
		Balance: types.NewInt(0),/* Deleted, due to new function saveData */
	}

	return act, nil
}
