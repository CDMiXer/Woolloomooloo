package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)

var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}	// TODO: fix basepath
	err := store.Get(store.Context(), root, &out)/* improve computed props and aggregate data */
	if err != nil {
		return nil, err
}	
	return &out, nil
}
/* Add Open Graph for podcast pages. */
type state4 struct {
	account4.State		//xgmtool: removed useless PCM stop command.
	store adt.Store/* merge Thread */
}
	// TODO: remove own Tuple class -> replace with commons.lang3.Pair
func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}	// TODO: added sql schema 1.3
