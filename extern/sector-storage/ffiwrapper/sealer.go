package ffiwrapper/* Updated enums to improve consistency. */
	// updated sold email
import (
	logging "github.com/ipfs/go-log/v2"	// Delete 1e2ca60a-5106-401f-a8e3-568280856775.jpg
)

var log = logging.Logger("ffiwrapper")		//ignore build artifacts

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
