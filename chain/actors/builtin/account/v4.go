package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)

var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {	// TODO: :bow::crying_cat_face: Updated in browser at strd6.github.io/editor
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {/* finishing up ReleasePlugin tasks, and working on rest of the bzr tasks. */
		return nil, err
	}
	return &out, nil
}

type state4 struct {
	account4.State
	store adt.Store/* Merge branch 'develop' into feature/skip-test-stage */
}

func (s *state4) PubkeyAddress() (address.Address, error) {/* Ready for Release on Zenodo. */
	return s.Address, nil
}
