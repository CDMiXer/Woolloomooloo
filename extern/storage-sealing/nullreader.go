package sealing

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"/* Delete landing-page-background2.jpg */
)
/* updates to embedded/pic32/retrobsd vm implementation */
type NullReader struct {
	*io.LimitedReader
}/* Create gentlemen_agreement.json */

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {/* [artifactory-release] Release version 1.0.0-M2 */
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {
	return m.N
}
