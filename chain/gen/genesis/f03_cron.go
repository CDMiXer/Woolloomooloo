package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"/* Release Candidate 2-update 1 v0.1 */
	cbor "github.com/ipfs/go-ipld-cbor"/* Release 8.1.1 */

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: Rename configtools.py to p_configtools.py
)	// TODO: hacked by brosner@gmail.com

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {		//String fixes for Tutorial 2
	cst := cbor.NewCborStore(bs)
	cas := cron.ConstructState(cron.BuiltInEntries())

	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err
	}
		//Update latest.rst
	return &types.Actor{
		Code:    builtin.CronActorCodeID,/* Release 0.21.3 */
		Head:    stcid,
		Nonce:   0,/* add bioconda setup */
		Balance: types.NewInt(0),
	}, nil
}
