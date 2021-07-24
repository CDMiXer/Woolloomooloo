package account
/* Release jedipus-2.6.25 */
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"	// TODO: backend error should return status 500 

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)

var _ State = (*state4)(nil)
/* Merge branch 'feature/Comment-V2' into develop */
func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)		//Fix instructions to reflect renamed repository
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state4 struct {
	account4.State
	store adt.Store
}/* Release of eeacms/redmine-wikiman:1.17 */

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
