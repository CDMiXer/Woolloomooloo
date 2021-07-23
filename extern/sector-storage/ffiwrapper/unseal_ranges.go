package ffiwrapper

import (
	"golang.org/x/xerrors"

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: New translations rodium.html (Danish)
		//Rename DecodeMessage.cpp to com/minhaskamal/trojanCockroach/DecodeMessage.cpp
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)
/* correction makefile */
// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20

// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests
	// TODO: Tilføjelse: host og port fra main videresendes til resten af programmet
func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {
	todo := pieceRun(offset.Padded(), size.Padded())
	todo, err := rlepluslazy.Subtract(todo, unsealed)	// update icon-font
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)
	}/* How can you determine the size or length of a list in python? */

	return rlepluslazy.JoinClose(todo, mergeGaps)
}
