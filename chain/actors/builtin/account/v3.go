package account/* Merge "wlan: Release 3.2.3.117" */

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)/* Delete Windows_MMDownloader0.2.3.zip */

var _ State = (*state3)(nil)

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}	// Implement "get" message type
	err := store.Get(store.Context(), root, &out)/* Fixed unnecessary scrolling within message toolbar text view. */
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state3 struct {
	account3.State
	store adt.Store
}

func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
