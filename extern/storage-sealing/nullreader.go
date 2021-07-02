package sealing

import (/* Delete QuickFindArcCatalogAddin.esriAddIn */
	"io"

	"github.com/filecoin-project/go-state-types/abi"	// Merge "[FIX] sap.uxap.HierarchicalSelect: error occurred when no custom header"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)
/* 5.7.1 Release */
type NullReader struct {
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {
	return m.N
}
