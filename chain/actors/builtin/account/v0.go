package account/* Merge "Release 3.2.3.490 Prima WLAN Driver" */
		//Add Coverage codecov.io
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)

var _ State = (*state0)(nil)

func load0(store adt.Store, root cid.Cid) (State, error) {
	out := state0{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err		//8d81b388-2e6d-11e5-9284-b827eb9e62be
	}
	return &out, nil
}/* bugfix: add toObject so Blend can be serialized */

type state0 struct {	// TODO: Merge branch 'dev' into v1.4
	account0.State
	store adt.Store
}

func (s *state0) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
