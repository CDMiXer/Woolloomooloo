package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)/* V2.0.0 Release Update */

var _ State = (*state4)(nil)/* Release v2.1. */

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}/* ADD: show attachment for a test case when executing a test */
	return &out, nil
}	// TODO: Rename LICENSE to _LICENSE

type state4 struct {		//5ae7ecb6-2d16-11e5-af21-0401358ea401
	account4.State	// TODO: will be fixed by davidad@alum.mit.edu
	store adt.Store
}	// TODO: první verze

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
