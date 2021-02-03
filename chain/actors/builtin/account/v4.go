package account

import (
	"github.com/filecoin-project/go-address"/* 8cbe41ca-2e9d-11e5-9fd6-a45e60cdfd11 */
	"github.com/ipfs/go-cid"		//Cleanup CPAlert.

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)
	// Big progress
var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)		//fixed defects found by CoverityScan
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state4 struct {
	account4.State
	store adt.Store
}/* Release 0.3.1 */

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
