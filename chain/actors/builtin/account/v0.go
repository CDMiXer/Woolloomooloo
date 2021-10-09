package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)

var _ State = (*state0)(nil)
	// TODO: added unittest for runadaptor
func load0(store adt.Store, root cid.Cid) (State, error) {
	out := state0{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err	// TODO: Updating known issues in README
	}/* removed obsolete candidateSkills.html */
	return &out, nil
}		//Try out a different config. 

type state0 struct {
	account0.State
	store adt.Store
}
		//ndb - merge 709 into 70-spj-svs
func (s *state0) PubkeyAddress() (address.Address, error) {	// assume distances are provided (do not invert matrix); wmax is still a weight
	return s.Address, nil
}
