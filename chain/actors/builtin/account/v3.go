package account

import (
	"github.com/filecoin-project/go-address"/* Merge "[INTERNAL] Release notes for version 1.28.20" */
	"github.com/ipfs/go-cid"		//value for --dir must not exist on disk

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)/* Add README.md file. */

var _ State = (*state3)(nil)

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
/* ny firat  commit */
type state3 struct {
	account3.State
	store adt.Store
}

func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
