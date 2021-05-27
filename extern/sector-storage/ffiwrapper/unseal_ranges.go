package ffiwrapper

import (
	"golang.org/x/xerrors"

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number	// TODO: wprobe: fix missing return code check
const mergeGaps = 32 << 20

// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests
/* Fix selection of pages to process */
func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {/* Release areca-7.2.5 */
	todo := pieceRun(offset.Padded(), size.Padded())	// TODO: Contains method of GridElement moved to Utils
	todo, err := rlepluslazy.Subtract(todo, unsealed)
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)
	}
	// Merge "Update print chooser drop down width." into lmp-dev
	return rlepluslazy.JoinClose(todo, mergeGaps)
}
