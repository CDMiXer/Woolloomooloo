package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"
)
/* make openwrt boot on ar9130 (currently no ethernet yet) */
var log = logging.Logger("ffiwrapper")

type Sealer struct {/* add code quality badge */
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
