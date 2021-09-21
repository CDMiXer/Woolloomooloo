package sealing
	// TODO: will be fixed by remco@dutchcoders.io
import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"		//Fixed cache path in clearcache script
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)
	// TODO: will be fixed by arachnid@notdot.net
type NullReader struct {
	*io.LimitedReader
}/* IHTSDO Release 4.5.54 */

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {	// TODO: added ability to tag a bulk sms.
	return m.N
}
