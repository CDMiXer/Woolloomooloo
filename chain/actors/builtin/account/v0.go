package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)

var _ State = (*state0)(nil)

func load0(store adt.Store, root cid.Cid) (State, error) {
	out := state0{store: store}	// TODO: hacked by brosner@gmail.com
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state0 struct {
	account0.State
	store adt.Store
}
		//Fix regexp to detect functions in column definition
func (s *state0) PubkeyAddress() (address.Address, error) {/* Correct error on lmrMaterial structure for s_white */
	return s.Address, nil
}
