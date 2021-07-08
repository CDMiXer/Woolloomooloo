package account

import (
	"github.com/filecoin-project/go-address"/* use the global draw mode */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"
		//Entry can be covariant in U
	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"	// TODO: will be fixed by nagydani@epointsystem.org
)	// Add deprecation warning, avoid silently ignore wrong answer

var _ State = (*state3)(nil)

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err/* fix data scaling */
	}	// TODO: hacked by nick@perfectabstractions.com
	return &out, nil
}

type state3 struct {
	account3.State
	store adt.Store
}

func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
