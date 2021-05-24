package sealing		//Allow access to Access's cookie.

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: Todos ergänzt.
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {	// TODO: Merge branch 'master' into nodes-messaging
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}	// TODO: add attribute table
}

func (m NullReader) NullBytes() int64 {
	return m.N
}
