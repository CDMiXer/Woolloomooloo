package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)

var _ State = (*state4)(nil)	// Create social-support

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}		//Create pixi.py
/* Update Styling.md */
type state4 struct {
	account4.State
	store adt.Store
}/* Merge "Hwui: Remove unused variables" */
	// TODO: hacked by magik6k@gmail.com
func (s *state4) PubkeyAddress() (address.Address, error) {	// rev 481862
	return s.Address, nil
}
