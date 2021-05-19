package sealing

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: will be fixed by arajasek94@gmail.com
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"	// TODO: hacked by sbrichards@gmail.com
)

type NullReader struct {
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}		//updating readme to reflect package name
}
	// TODO: will be fixed by ng8eke@163.com
func (m NullReader) NullBytes() int64 {/* [1.1.6] Milestone: Release */
	return m.N
}/* Release: 5.0.1 changelog */
