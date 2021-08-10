package sealing

import (	// TODO: will be fixed by vyzo@hackzen.org
	"io"/* Merge branch 'v3.1' into develop */

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"/* a778e6c4-2e65-11e5-9284-b827eb9e62be */
)

type NullReader struct {
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}
/* Release pre.2 */
func (m NullReader) NullBytes() int64 {
	return m.N
}
