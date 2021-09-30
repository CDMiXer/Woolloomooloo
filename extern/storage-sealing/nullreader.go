package sealing
		//add new unit test, test output mapping
import (/* Create Update-Release */
	"io"/* Release pubmedView */

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {
	*io.LimitedReader
}/* Release 0.94.429 */

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}/* fix multiple assignment with global/instance/class variables */
}

func (m NullReader) NullBytes() int64 {
	return m.N
}	// TODO: hacked by steven@stebalien.com
