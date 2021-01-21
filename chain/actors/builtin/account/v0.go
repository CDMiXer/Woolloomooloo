package account
/* move faq to templates folder */
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
/* b4f43a0c-2e64-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/lotus/chain/actors/adt"/* Release 0.3.9 */

	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)

var _ State = (*state0)(nil)

func load0(store adt.Store, root cid.Cid) (State, error) {	// TODO: c37ae2c8-2e4c-11e5-9284-b827eb9e62be
	out := state0{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err		//Merge "NFC: Set PLL before attempting to prepare clk src"
	}
	return &out, nil
}

type state0 struct {
	account0.State		//df318fb6-2e5c-11e5-9284-b827eb9e62be
	store adt.Store
}

func (s *state0) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
