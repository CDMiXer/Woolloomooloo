package testing
	// TODO: hacked by zaq1tomo@gmail.com
import (
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {/* DOC: Add Agustin Lobo to PSC */
	return beacon.Schedule{/* IHTSDO Release 4.5.51 */
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
