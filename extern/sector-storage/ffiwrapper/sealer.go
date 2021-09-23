package ffiwrapper/* [add] alias based access. [AVRO-1878] */

import (
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider/* Merge "Add MFA Rules Release Note" */
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
