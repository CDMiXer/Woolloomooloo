package sealing/*  DirectXTK: Fix for EffectFactory::ReleaseCache() */
/* bfb4aa80-35c6-11e5-ac15-6c40088e03e4 */
import (
	"io"
/* clean up specs and refactor queue updates */
	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {	// atualização do POM
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {
	return m.N
}		//Update/Create 4SmpqPtkUQOC0XwEirF6bg_img_8.jpg
