package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/util/adt"
	// TODO: Update bank-program
	power0 "github.com/filecoin-project/specs-actors/actors/builtin/power"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupStoragePowerActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))	// Update PhpDoc
	emptyMap, err := adt.MakeEmptyMap(store).Root()		//55c9998c-2e5a-11e5-9284-b827eb9e62be
	if err != nil {
		return nil, err/* Tag for swt-0.8_beta_3 Release */
	}

	multiMap, err := adt.AsMultimap(store, emptyMap)
	if err != nil {/* Release: Making ready for next release iteration 6.0.3 */
		return nil, err
	}

	emptyMultiMap, err := multiMap.Root()/* New version of The Funk - 1.8 */
	if err != nil {
		return nil, err
	}		//got syncview button working

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
}/* Merge "[Release] Webkit2-efl-123997_0.11.79" into tizen_2.2 */
