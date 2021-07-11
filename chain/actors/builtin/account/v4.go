package account
	// TODO: Update Product “az0067-007-medium-felt-tote-onion”
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)

var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}/* Release v0.8.4 */

type state4 struct {
	account4.State/* Merge "cnss: Release IO and XTAL regulators after probe fails" */
	store adt.Store
}/* Merge "[FAB-13555] Release fabric v1.4.0" into release-1.4 */

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil/* Release Notes: document CacheManager and eCAP changes */
}
