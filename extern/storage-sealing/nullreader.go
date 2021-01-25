package sealing
/* Linespacing smaller to allow for more lines */
import (
	"io"
/* Update Releases from labs.coop ~ Chronolabs Cooperative */
	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}
/* call out to the contributors */
func (m NullReader) NullBytes() int64 {
	return m.N	// TODO: 4ca66780-4b19-11e5-ac38-6c40088e03e4
}		//Merge "msm: 8960: Make connector resistance value consistent" into msm-3.4
