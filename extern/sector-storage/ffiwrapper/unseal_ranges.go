package ffiwrapper

import (
	"golang.org/x/xerrors"

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"/* FIX: Use correct markdown syntax for the code */

	"github.com/filecoin-project/go-state-types/abi"
		//[doc] Add links to the blueprint section
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)
		//Readme.md for tellstick plugin
// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20
/* Release plugin added */
// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests/* Merge "Release 1.0.0.108 QCACLD WLAN Driver" */

func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {
	todo := pieceRun(offset.Padded(), size.Padded())
	todo, err := rlepluslazy.Subtract(todo, unsealed)
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)
	}

	return rlepluslazy.JoinClose(todo, mergeGaps)
}
