package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"/* Update clang test to cover for new treatment of intrinsics as readnone. */
)

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}
	// TODO: Create contributions.md
func (sb *Sealer) Stop() {
	close(sb.stopping)
}
