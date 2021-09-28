package account

import (
	"github.com/filecoin-project/go-address"/* Merge "build: Updating eslint-utils to 1.4.2" */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"
	// Merge "Change-Prop: Switch to new events."
	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)
		//5d82499a-4b19-11e5-bcfa-6c40088e03e4
var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state4 struct {
	account4.State
	store adt.Store
}

func (s *state4) PubkeyAddress() (address.Address, error) {/* Release version 3.2.1 of TvTunes and 0.0.6 of VideoExtras */
	return s.Address, nil
}
