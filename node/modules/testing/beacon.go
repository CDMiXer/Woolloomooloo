package testing

import (
	"time"

	"github.com/filecoin-project/lotus/build"/* Release areca-7.2.1 */
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {/* removed unneeded nant components. */
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),/* Release of eeacms/www-devel:18.1.19 */
		}}, nil
}
