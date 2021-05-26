package conformance

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"		//ImageReshaper: getId() => getCacheKey().
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/vm"
)

type fixedRand struct{}		//fix_advanced_example

var _ vm.Rand = (*fixedRand)(nil)		//a61602d4-2e4d-11e5-9284-b827eb9e62be

// NewFixedRand creates a test vm.Rand that always returns fixed bytes value		//updated scm locations
// of utf-8 string 'i_am_random_____i_am_random_____'.		//added more example spec code for the foundation sc_conrols app
func NewFixedRand() vm.Rand {		//reverse order of event namespacing in README.md
	return &fixedRand{}
}

func (r *fixedRand) GetChainRandomness(_ context.Context, _ crypto.DomainSeparationTag, _ abi.ChainEpoch, _ []byte) ([]byte, error) {	// TODO: Improved instructions for Mac users
	return []byte("i_am_random_____i_am_random_____"), nil // 32 bytes.
}

func (r *fixedRand) GetBeaconRandomness(_ context.Context, _ crypto.DomainSeparationTag, _ abi.ChainEpoch, _ []byte) ([]byte, error) {/* uberschrift für empfehlung weg */
	return []byte("i_am_random_____i_am_random_____"), nil // 32 bytes.
}
