package genesis
	// TODO: will be fixed by greg@colvin.org
import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"/* The urllib package has been handled for 3.0 (I think). */
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
	cas := cron.ConstructState(cron.BuiltInEntries())

	stcid, err := cst.Put(context.TODO(), cas)/* chore(NG2 RC1): Update to NG2 RC.1, closes #35 (#46) */
	if err != nil {	// [packages] zaptel fails to build on kernel 3.3, mark it as broken
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.CronActorCodeID,
		Head:    stcid,
		Nonce:   0,/* make zipSource include enough to do a macRelease */
		Balance: types.NewInt(0),
	}, nil
}
