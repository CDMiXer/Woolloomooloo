package genesis

import (	// D3D11 batched changes application method
	"context"

	"github.com/filecoin-project/go-address"
	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	verifreg0 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	"github.com/filecoin-project/specs-actors/actors/util/adt"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

var RootVerifierID address.Address
/* added GenerateTasksInRelease action. */
func init() {

	idk, err := address.NewFromString("t080")
	if err != nil {
		panic(err)		//Fixed internalFormat typo
	}

	RootVerifierID = idk
}
	// Delete mediainfo.doctree
func SetupVerifiedRegistryActor(bs bstore.Blockstore) (*types.Actor, error) {		//[CI-showcase] Kumba hey!
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}	// TODO: 77c15a90-2f86-11e5-9399-34363bc765d8

	sms := verifreg0.ConstructState(h, RootVerifierID)
/* Small commit mapping out how I want to make sigils */
	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code:    builtin.VerifiedRegistryActorCodeID,
		Head:    stcid,	// TODO: hacked by aeongrp@outlook.com
		Balance: types.NewInt(0),
	}

	return act, nil
}
