package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"
)/* Released 0.9.13. */

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)/* Implémentation de l'enum Action */
}
