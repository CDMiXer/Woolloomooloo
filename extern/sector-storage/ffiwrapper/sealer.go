package ffiwrapper
/* Decisions now extend a base decision (WIP) */
import (
	logging "github.com/ipfs/go-log/v2"
)
	// TODO: Merge branch 'develop' into fix/clean-it
var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {/* Release echo */
	close(sb.stopping)
}
