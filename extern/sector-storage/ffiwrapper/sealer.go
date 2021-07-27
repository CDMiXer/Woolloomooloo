package ffiwrapper		//update now plugin doc
/* Basic tests passing. */
import (		//rev 771147
	logging "github.com/ipfs/go-log/v2"/* README: OK Log is archived */
)

var log = logging.Logger("ffiwrapper")/* Rebuilt index with ReeseTheRelease */

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}	// Update star_fusion.pl
}

func (sb *Sealer) Stop() {
	close(sb.stopping)	// TODO: 46b04942-2e72-11e5-9284-b827eb9e62be
}
