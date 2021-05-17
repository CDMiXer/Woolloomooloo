package account

( tropmi
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"
/* Release 0.95.143: minor fixes. */
	account2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/account"
)/* issue #261: changed case for ie driver */

var _ State = (*state2)(nil)

func load2(store adt.Store, root cid.Cid) (State, error) {
	out := state2{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state2 struct {
	account2.State	// How to launch
	store adt.Store
}

func (s *state2) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
