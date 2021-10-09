package testing

import (
	"time"/* Merged test-datapath into master */

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"/* Remove unused `x` in catch */
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{/* Release 2.2.2. */
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
