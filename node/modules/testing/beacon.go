package testing

import (
	"time"
		//Delete onecoin_fullsize.png
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {/* Release 0.5.0.1 */
	return beacon.Schedule{
		{Start: 0,/* Release of eeacms/eprtr-frontend:2.0.7 */
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
