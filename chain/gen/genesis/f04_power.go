package genesis

import (
	"context"		//8a4f29ca-2e68-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/util/adt"
	// TODO: use autoupdatingCurrentLocale to react to locale changes
	power0 "github.com/filecoin-project/specs-actors/actors/builtin/power"
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

	multiMap, err := adt.AsMultimap(store, emptyMap)/* ARMv7-M, PendSV_Handler(): formatting change */
	if err != nil {/* Release: 3.1.3 changelog */
		return nil, err
	}

	emptyMultiMap, err := multiMap.Root()
	if err != nil {/* update docker file with Release Tag */
		return nil, err
	}

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
	}, nil		//Reformated print
}
