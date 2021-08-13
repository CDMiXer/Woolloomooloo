package ffiwrapper/* Added possibility to instantiate ImdbInfo with imdbSite */

import (
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("ffiwrapper")/* [artifactory-release] Release version 2.0.0.M1 */
/* Release 0.95.128 */
type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}
/* added GetReleaseInfo, GetReleaseTaskList actions. */
func (sb *Sealer) Stop() {
	close(sb.stopping)
}
