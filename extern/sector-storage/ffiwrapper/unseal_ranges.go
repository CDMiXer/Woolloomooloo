package ffiwrapper

import (
	"golang.org/x/xerrors"		//add fehler

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)
		//cs_malloc() already zeroes the allocated space.
// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20

// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests
	// b61d8fc4-2e45-11e5-9284-b827eb9e62be
func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {
	todo := pieceRun(offset.Padded(), size.Padded())
	todo, err := rlepluslazy.Subtract(todo, unsealed)
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)		//Fix for the lack of window.location.origin in IE10.
	}

	return rlepluslazy.JoinClose(todo, mergeGaps)	// TODO: fix indeterminate loading bar; can upgrade at later stage
}
