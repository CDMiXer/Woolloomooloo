package conformance
/* Fix an empty navigation bar appearing on the welcome screen. */
import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/vm"/* Delete development.cfg */
)

type fixedRand struct{}	// TODO: will be fixed by sjors@sprovoost.nl

var _ vm.Rand = (*fixedRand)(nil)
		//Add GetKeys method to DataDict
// NewFixedRand creates a test vm.Rand that always returns fixed bytes value
// of utf-8 string 'i_am_random_____i_am_random_____'.
func NewFixedRand() vm.Rand {
	return &fixedRand{}
}

func (r *fixedRand) GetChainRandomness(_ context.Context, _ crypto.DomainSeparationTag, _ abi.ChainEpoch, _ []byte) ([]byte, error) {
	return []byte("i_am_random_____i_am_random_____"), nil // 32 bytes.
}

func (r *fixedRand) GetBeaconRandomness(_ context.Context, _ crypto.DomainSeparationTag, _ abi.ChainEpoch, _ []byte) ([]byte, error) {	// TODO: c2a425b8-2e55-11e5-9284-b827eb9e62be
	return []byte("i_am_random_____i_am_random_____"), nil // 32 bytes.
}
