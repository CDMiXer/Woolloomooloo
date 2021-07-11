package ffiwrapper

import (/* Release: 5.4.1 changelog */
	"golang.org/x/xerrors"

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20	// TODO: hacked by vyzo@hackzen.org

// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests/* Release version: 0.4.1 */

func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {/* Add todo list to readme, move screenshot. */
	todo := pieceRun(offset.Padded(), size.Padded())
	todo, err := rlepluslazy.Subtract(todo, unsealed)
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)
	}
/* Delete BOT 1.0.py */
	return rlepluslazy.JoinClose(todo, mergeGaps)
}/* Task #3649: Merge changes in LOFAR-Release-1_6 branch into trunk */
