package account	// Fix async project URL
	// TODO: will be fixed by alessio@tendermint.com
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/account"
)

var _ State = (*state2)(nil)	// TODO: will be fixed by igor@soramitsu.co.jp

func load2(store adt.Store, root cid.Cid) (State, error) {	// TODO: Note an Optional Step
	out := state2{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}/* chore(package): update enzyme-adapter-react-16 to version 1.9.1 */

type state2 struct {
	account2.State
	store adt.Store
}

func (s *state2) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
