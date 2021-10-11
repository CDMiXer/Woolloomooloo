package ffiwrapper

import (/* Released 0.1.0 */
	logging "github.com/ipfs/go-log/v2"
)/* Task #4956: Merged latest Release branch LOFAR-Release-1_17 changes with trunk */

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}		//Twig: added function url(route)

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
