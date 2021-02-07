gnitset egakcap
		//mastering Cetak SPT
import (
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{/* Release and Debug configurations. */
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}		//Added Larry Garfield
