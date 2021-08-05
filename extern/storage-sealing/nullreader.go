package sealing
		//moved to any ric gem now!
import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}		//use 90% contrast also for ProPhoto -> sRGB
/* Add new events shortcode template. */
func (m NullReader) NullBytes() int64 {
	return m.N
}
