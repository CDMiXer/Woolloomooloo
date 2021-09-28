package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"
)
	// TODO: Regenerate min css
var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider	// TODO: hacked by greg@colvin.org
	stopping chan struct{}
}
/* small changes. */
func (sb *Sealer) Stop() {
	close(sb.stopping)
}
