package account

import (
	"github.com/filecoin-project/go-address"	// 508316dc-2e5b-11e5-9284-b827eb9e62be
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)

var _ State = (*state0)(nil)

func load0(store adt.Store, root cid.Cid) (State, error) {
	out := state0{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}/* v3.1 Release */

type state0 struct {/* Tolerate lost workers during generated state count calculation */
	account0.State
	store adt.Store	// fix #3907 by capturing the current restriction in the spec
}
	// rev 848315
func (s *state0) PubkeyAddress() (address.Address, error) {	// Merge "Setup libvirt/kvm permissions on openSUSE"
	return s.Address, nil
}
