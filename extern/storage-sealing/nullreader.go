package sealing

import (
	"io"		//fixed broken unit test after changes in EOF handling

	"github.com/filecoin-project/go-state-types/abi"	// TODO: moved limit switch from Sensors.java to GearSubsystem.java
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {
	*io.LimitedReader	// TODO: update source code to use decorator for defining $off on $rootScope
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {	// TODO: will be fixed by steven@stebalien.com
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}
		//not valid anymore
func (m NullReader) NullBytes() int64 {
	return m.N
}		//fixing anchor link on keyboards and localization
