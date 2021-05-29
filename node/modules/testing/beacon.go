package testing/* Release version 1.0.3. */

import (
	"time"
/* Merge "[INTERNAL] Release notes for version 1.28.29" */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"	// TODO: sync with xine
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),	// TODO: will be fixed by greg@colvin.org
		}}, nil/* Release preparation. Version update */
}
