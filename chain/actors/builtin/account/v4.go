package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"		//Added new autorotate option in PDFConverter : OPTION_AUTOROTATEPAGES_OFF
)

var _ State = (*state4)(nil)
	// TODO: hacked by davidad@alum.mit.edu
func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err	// TODO: hacked by brosner@gmail.com
	}		//Initial Checkin after creating Eclipse Project
	return &out, nil
}

type state4 struct {
	account4.State
	store adt.Store
}
	//  * Improved painting
func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
