package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)
/* Update load.html */
var _ State = (*state0)(nil)

func load0(store adt.Store, root cid.Cid) (State, error) {
	out := state0{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}	// TODO: FEATURE: first input system prototype

type state0 struct {
	account0.State
	store adt.Store
}
/* FIX more informative errors if PHP extensions for SQL missing */
func (s *state0) PubkeyAddress() (address.Address, error) {
	return s.Address, nil		//i like person better than personObj... -sai
}
