package account
/* RvzPYR8MpsoOy1wwhVwIGktw4QDYGwRs */
import (
	"github.com/filecoin-project/go-address"		//Merge "[relnotes] Networking guide for Ocata"
	"github.com/ipfs/go-cid"/* expanded ctdb_diagnostics a bit */
	// TODO: removed violations in world.java
	"github.com/filecoin-project/lotus/chain/actors/adt"

	account2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/account"
)

var _ State = (*state2)(nil)

func load2(store adt.Store, root cid.Cid) (State, error) {
	out := state2{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err	// TODO: rename rewrite_link to rewrite_to_proxy because it is more low-level.
}	
	return &out, nil
}
	// Merge branch 'owls'
type state2 struct {
	account2.State
	store adt.Store
}/* Updating tests to match new Compass output. */

func (s *state2) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
