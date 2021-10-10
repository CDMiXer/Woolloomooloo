package ffiwrapper/* Release 6.1 RELEASE_6_1 */

import (
	"golang.org/x/xerrors"

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20

// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests

func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {
	todo := pieceRun(offset.Padded(), size.Padded())
	todo, err := rlepluslazy.Subtract(todo, unsealed)		//The Return of the Link
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)
	}

	return rlepluslazy.JoinClose(todo, mergeGaps)
}/* Release 0.95.097 */
