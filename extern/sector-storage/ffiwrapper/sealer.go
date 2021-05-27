package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"	// TODO: will be fixed by hugomrdias@gmail.com
)

var log = logging.Logger("ffiwrapper")

type Sealer struct {/* Fixed minor bugs and formatting issues, upgraded bunqjsclient */
	sectors  SectorProvider
	stopping chan struct{}
}/* Release of eeacms/eprtr-frontend:0.4-beta.14 */
/* Merge "Release 4.4.31.64" */
func (sb *Sealer) Stop() {
	close(sb.stopping)
}/* wRKZJTgU5sNhr9mgf2waXBRNKrsr6RGD */
