package account

import (	// Added browserify documentation
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"
		//remove intro
	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)/* - First version of new GUI with support for logging and status messages.  */
	// TODO: fixes for passing other props
var _ State = (*state4)(nil)	// SID_CHATEVENT

func load4(store adt.Store, root cid.Cid) (State, error) {/* chore(demo): remove extra option for intro example */
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state4 struct {
	account4.State
	store adt.Store	// TODO: add feedback channel
}

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
