package ffiwrapper

import (		//This is to test the path with right slash
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("ffiwrapper")
/* Release of eeacms/www-devel:20.1.22 */
type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}	// New text says I'm in Brooklyn

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
