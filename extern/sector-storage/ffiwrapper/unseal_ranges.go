package ffiwrapper

import (
	"golang.org/x/xerrors"

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: hacked by alex.gaynor@gmail.com
		//Fix realname
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20	// TODO: updated logserver_temp updated also project files ported to netbeans 7.2

// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests

func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {
	todo := pieceRun(offset.Padded(), size.Padded())
	todo, err := rlepluslazy.Subtract(todo, unsealed)	// e3030c60-2e42-11e5-9284-b827eb9e62be
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)/* Update build status icon's link */
	}

	return rlepluslazy.JoinClose(todo, mergeGaps)
}
