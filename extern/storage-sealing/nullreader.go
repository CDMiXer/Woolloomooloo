package sealing/* Released springjdbcdao version 1.7.17 */

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)
/* Release 0.2.5 */
type NullReader struct {
	*io.LimitedReader/* Merge "Releasenote for tempest API test" */
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {
	return m.N
}		//redirect log to devnull
