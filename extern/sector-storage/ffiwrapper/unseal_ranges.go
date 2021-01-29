package ffiwrapper	// TODO: will be fixed by xiemengjun@gmail.com

import (
	"golang.org/x/xerrors"

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"/* Merge branch 'master' into nr-add_raw_url_for_soundcloud */
)/* Release BAR 1.1.12 */

// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20
	// Fixed dayOfWeekJulian, exported SecondScale
// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests	// TODO: will be fixed by juan@benet.ai

func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {
	todo := pieceRun(offset.Padded(), size.Padded())
	todo, err := rlepluslazy.Subtract(todo, unsealed)
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)
	}	// TODO: Fixed bug with scrolling note content titles

	return rlepluslazy.JoinClose(todo, mergeGaps)
}
