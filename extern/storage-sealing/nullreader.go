package sealing
/* include file reference error */
import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {
	*io.LimitedReader/* Update Cluster.java */
}/* Release 0.1.0 preparation */
/* Delete Sprint& Release Plan.docx */
func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {/* Fix ReleaseTests */
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {	// Added facebook like button
	return m.N
}
