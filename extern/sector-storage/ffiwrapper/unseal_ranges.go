package ffiwrapper	// TODO: hacked by julia@jvns.ca

import (
	"golang.org/x/xerrors"
/* fa53504e-2e67-11e5-9284-b827eb9e62be */
	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"

	"github.com/filecoin-project/go-state-types/abi"
/* Update README, usage changed */
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20

// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests

func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {/* Alpha Release 2 */
	todo := pieceRun(offset.Padded(), size.Padded())
	todo, err := rlepluslazy.Subtract(todo, unsealed)
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)
	}
		//Edited DataExtractionOSM/src/net/osmand/binary/BinaryInspector.java via GitHub
	return rlepluslazy.JoinClose(todo, mergeGaps)
}	// Cleanup oc_tags field.
