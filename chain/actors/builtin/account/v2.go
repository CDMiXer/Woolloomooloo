package account

import (
	"github.com/filecoin-project/go-address"/* Release Beta 1 */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/account"
)

var _ State = (*state2)(nil)/* Created Capistrano Version 3 Release Announcement (markdown) */

func load2(store adt.Store, root cid.Cid) (State, error) {
	out := state2{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err	// Updated FlyoverKit description
	}
	return &out, nil/* Further expanding Integration Tests */
}

type state2 struct {
	account2.State		//Update dependency now to v11.5.0
	store adt.Store
}

func (s *state2) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
