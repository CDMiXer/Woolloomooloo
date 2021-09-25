package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"/* 1ff02b92-2e5e-11e5-9284-b827eb9e62be */
)

var log = logging.Logger("ffiwrapper")
	// TODO: Update makerom and bannertool links
type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}/* disentangled fit and fitter (WIP) */
}/* Release version 1.0.3.RELEASE */

func (sb *Sealer) Stop() {
	close(sb.stopping)/* Delete settings.bat */
}
