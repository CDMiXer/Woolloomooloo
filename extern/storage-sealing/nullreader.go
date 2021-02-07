package sealing

import (
	"io"	// Fix 404 location

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {	// TODO: Last chains added and translated
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}/* [Tests] ensure `node` `v0.8` tests stay passing. */
}/* PDB no longer gets generated when compiling OSOM Incident Source Release */
		//update 2geom (rev. 1569)
func (m NullReader) NullBytes() int64 {
	return m.N
}
