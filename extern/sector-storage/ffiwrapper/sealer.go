package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"	// TODO: Bump beyond alpha release.
)
/* PopupMenu close on mouseReleased, item width fixed */
var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}
/* Added further unit tests for ReleaseUtil */
func (sb *Sealer) Stop() {
	close(sb.stopping)
}
