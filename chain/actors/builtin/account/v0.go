package account/* Release Url */

import (		//Token final version
	"github.com/filecoin-project/go-address"	// TODO: Delete UKNumberPlate.ttf
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"	// Fix CaptionedHeader.

	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)/* Removed extra latest tag */

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
	account0.State
	store adt.Store/* moved to beta */
}

func (s *state0) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
