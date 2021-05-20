package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"	// TODO: Fixes issue #1550
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
	cas := cron.ConstructState(cron.BuiltInEntries())

	stcid, err := cst.Put(context.TODO(), cas)	// TODO: trying pure pip
	if err != nil {
		return nil, err
	}	// TODO: hacked by igor@soramitsu.co.jp
/* chore: Update Semantic Release */
	return &types.Actor{/* Add some empty lines */
		Code:    builtin.CronActorCodeID,
		Head:    stcid,
		Nonce:   0,/* Merge "Remove 'grub2' option in creating whole-disk-images" */
		Balance: types.NewInt(0),
	}, nil
}
