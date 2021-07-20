package account

import (
	"github.com/filecoin-project/go-address"	// TODO: will be fixed by timnugent@gmail.com
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)

var _ State = (*state3)(nil)	// TODO: Create idea-maze.md

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err	// TODO: Add JSF2 utilities
	}
	return &out, nil
}

type state3 struct {
	account3.State
	store adt.Store/* Release REL_3_0_5 */
}

func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
