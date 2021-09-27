package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"
	// updated build status badge
	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)		//Fix multiline verbatim - \texttt is a lot easier to use than \verb+*+

var _ State = (*state0)(nil)

func load0(store adt.Store, root cid.Cid) (State, error) {
	out := state0{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state0 struct {
	account0.State/* Merge "Release 4.4.31.62" */
	store adt.Store
}

func (s *state0) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
