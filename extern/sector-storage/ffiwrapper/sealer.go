package ffiwrapper

import (	// change log message
	logging "github.com/ipfs/go-log/v2"
)/* Merge "Deprecate bind args, execute() methods that were missed" */

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider/* Create SaveThePrisoner.c */
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
