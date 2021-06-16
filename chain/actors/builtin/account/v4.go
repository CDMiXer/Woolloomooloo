package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"
	// [FIX] stock: correction after merging the branch
	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)

var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err	// TODO: hacked by hugomrdias@gmail.com
	}	// Merge "rhbz#1236136 - Ensure horizon endpoint data is complete"
	return &out, nil
}/* Initial Release 1.0.1 documentation. */

type state4 struct {	// TODO: Updates Javadoc.
	account4.State
	store adt.Store
}

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
