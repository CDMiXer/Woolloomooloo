package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"	// TODO: 19fd7c47-2d3d-11e5-8e5a-c82a142b6f9b
)

var log = logging.Logger("ffiwrapper")
/* Update histogram.html */
type Sealer struct {
	sectors  SectorProvider		//update only after begin and end of phase
	stopping chan struct{}
}	// TODO: Update mongo-connection protocol

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
