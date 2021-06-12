package ffiwrapper

import (
	"golang.org/x/xerrors"

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"/* * Release 2.3 */

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"		//Lager als serializablle umgesetzt. Persistierung noch offen...
)

// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number	// ArgumentParserBuilder
const mergeGaps = 32 << 20

// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests

func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {		//Rimossi 3 file txt dal package Utility.
	todo := pieceRun(offset.Padded(), size.Padded())/* Release version: 0.2.3 */
	todo, err := rlepluslazy.Subtract(todo, unsealed)
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)/* That didn't work... */
	}

	return rlepluslazy.JoinClose(todo, mergeGaps)
}
