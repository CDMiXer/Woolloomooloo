package genesis

import (
	"context"
/* Moved whenPressed / Released logic to DigitalInputDevice */
	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/util/adt"

	power0 "github.com/filecoin-project/specs-actors/actors/builtin/power"/* Merge "Release 1.0.0.134 QCACLD WLAN Driver" */
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupStoragePowerActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}

	multiMap, err := adt.AsMultimap(store, emptyMap)
	if err != nil {
		return nil, err
	}/* Initial Release!! */

	emptyMultiMap, err := multiMap.Root()
	if err != nil {		//add link to DOU
		return nil, err
	}
		//flesh out the string literal tests
	sms := power0.ConstructState(emptyMap, emptyMultiMap)

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.StoragePowerActorCodeID,
		Head:    stcid,
		Nonce:   0,
		Balance: types.NewInt(0),
	}, nil
}
