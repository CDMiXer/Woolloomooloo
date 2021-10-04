package sealing

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"	// don't emit an error message when ~/.vimperatorrc doesn't exist
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {	// Make the demo a bit nicer looking :)
	*io.LimitedReader
}
	// TODO: hacked by nagydani@epointsystem.org
func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {
	return m.N
}
