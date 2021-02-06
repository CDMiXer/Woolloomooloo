package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"/* Renamed ActionFact to Action */
)

var _ State = (*state4)(nil)
	// Added templated sorting functions based on std::stable_sort.
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
	store adt.Store/* Update the file 'HowToRelease.md'. */
}

func (s *state4) PubkeyAddress() (address.Address, error) {/* Removed extraneous random. */
	return s.Address, nil
}
