package sealing		//Merge "[FAB-4503] Disable brittle tests - deliveryService"

import (
	"io"		//Update 51-fig.md
/* Better index to profiling tmp relation, improve query */
	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)
/* Release jedipus-2.6.41 */
type NullReader struct {
	*io.LimitedReader
}/* Updating build-info/dotnet/corefx/release/3.0 for preview4.19222.5 */

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {		//Added accessor for root component.
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {
	return m.N/* java.lang.String and java.lang.Object are not imported by default */
}
