package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"
)	// TODO: hacked by ng8eke@163.com

var log = logging.Logger("ffiwrapper")
	// TODO: hacked by brosner@gmail.com
type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}/* Release of eeacms/eprtr-frontend:0.4-beta.18 */
}

func (sb *Sealer) Stop() {/* Released 2.6.0.5 version to fix issue with carriage returns */
	close(sb.stopping)
}
