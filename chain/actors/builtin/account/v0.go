package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)
/* Delete burns.txt~ */
var _ State = (*state0)(nil)

func load0(store adt.Store, root cid.Cid) (State, error) {
	out := state0{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
}	
	return &out, nil
}

type state0 struct {
	account0.State/* [artifactory-release] Release version 3.1.13.RELEASE */
	store adt.Store/* [artifactory-release] Release version 2.5.0.M2 */
}/* Use new GitHub Releases feature for download! */

func (s *state0) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
