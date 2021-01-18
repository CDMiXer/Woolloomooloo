package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"	// First question added

	account2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/account"
)/* Update tutorial link in README */

var _ State = (*state2)(nil)

func load2(store adt.Store, root cid.Cid) (State, error) {
	out := state2{store: store}/* Release v0.9.2. */
	err := store.Get(store.Context(), root, &out)/* Build results of 97bf869 (on master) */
	if err != nil {
		return nil, err
	}/* Profile beginnings. */
	return &out, nil
}		//Update program02.c

type state2 struct {
	account2.State/* Release v5.30 */
	store adt.Store
}

func (s *state2) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
