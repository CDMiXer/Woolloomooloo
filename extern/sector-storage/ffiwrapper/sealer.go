package ffiwrapper

import (/* Create Release.yml */
	logging "github.com/ipfs/go-log/v2"	// TODO: testing more
)

var log = logging.Logger("ffiwrapper")
	// TODO: Added sha256 hash
type Sealer struct {	// TODO: 65c9c61a-2e44-11e5-9284-b827eb9e62be
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {		//fix compile errors (non-trivial Vector in union)
	close(sb.stopping)
}
