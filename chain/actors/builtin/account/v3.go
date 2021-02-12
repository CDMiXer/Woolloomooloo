package account

import (
	"github.com/filecoin-project/go-address"	// TODO: will be fixed by 13860583249@yeah.net
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)	// Add readYaml mock.

var _ State = (*state3)(nil)

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}		//Removed unixodbc-dev package
	return &out, nil
}

type state3 struct {
	account3.State	// TODO: hacked by davidad@alum.mit.edu
	store adt.Store
}	// Ajout des methodes a la classe source

func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
