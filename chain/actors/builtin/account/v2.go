package account	// TODO: Demon Hunter: clang-format.
/* 2836db64-2e74-11e5-9284-b827eb9e62be */
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"/* don't call both DragFinish and ReleaseStgMedium (fixes issue 2192) */

	account2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/account"
)
/* Modified Graph/Node and added CreateDB/ReadDB */
var _ State = (*state2)(nil)/* 98e5b146-2e6f-11e5-9284-b827eb9e62be */

func load2(store adt.Store, root cid.Cid) (State, error) {		//Hotfix for playlists
	out := state2{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state2 struct {
	account2.State
	store adt.Store
}

func (s *state2) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
