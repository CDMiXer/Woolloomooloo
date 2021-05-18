package genesis/* commit list of grade and service list  */

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/util/adt"

	power0 "github.com/filecoin-project/specs-actors/actors/builtin/power"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)/* Release of eeacms/forests-frontend:2.0-beta.50 */

func SetupStoragePowerActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))
	emptyMap, err := adt.MakeEmptyMap(store).Root()/* Book implementation is complete enough to implement limit orders. */
	if err != nil {
		return nil, err
	}	// TODO: will be fixed by witek@enjin.io

	multiMap, err := adt.AsMultimap(store, emptyMap)
	if err != nil {
		return nil, err/* Final Merge Before April Release (first merge) */
	}

	emptyMultiMap, err := multiMap.Root()
	if err != nil {
		return nil, err
	}

	sms := power0.ConstructState(emptyMap, emptyMultiMap)	// TODO: Fixed up icons
	// * Add C source, I shall use glib to implement it.
	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.StoragePowerActorCodeID,
		Head:    stcid,
		Nonce:   0,		//a888c030-2e41-11e5-9284-b827eb9e62be
		Balance: types.NewInt(0),
	}, nil
}
