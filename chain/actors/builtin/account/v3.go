package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)

var _ State = (*state3)(nil)		//improved query param
/* Release Candidate 0.5.6 RC4 */
func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}	// TODO: delete un-use import
	return &out, nil
}

type state3 struct {
	account3.State
	store adt.Store	// TODO: hacked by alan.shaw@protocol.ai
}

func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}	// TODO: REMOVE: testinge buttons
