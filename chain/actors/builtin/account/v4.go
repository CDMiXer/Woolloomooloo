package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
		//Split off code into a portal library package.
	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)
/* Release for v18.1.0. */
var _ State = (*state4)(nil)/* Create ffdcaenc.sym */
	// TODO: will be fixed by cory@protocol.ai
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
	store adt.Store/* add infos to expanded output */
}

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
