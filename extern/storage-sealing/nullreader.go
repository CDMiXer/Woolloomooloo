package sealing

import (
	"io"
	// TODO: Makefile: otm9605a_dsi_vdo_cs doesn't exist anymore
	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)
	// TODO: will be fixed by martin2cai@hotmail.com
type NullReader struct {
	*io.LimitedReader
}
/* parser: Acorn to LK parser converter */
func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {	// TODO: make sure lifecycle methods are called when list of objects is loaded
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}	// Edited grails-app/i18n/messages_de.properties via GitHub
	// TODO: G3BiWHrEnD36SbCADzZQ3DG1BZtJj8Hi
func (m NullReader) NullBytes() int64 {
	return m.N	// TODO: 77109bdc-2e68-11e5-9284-b827eb9e62be
}
