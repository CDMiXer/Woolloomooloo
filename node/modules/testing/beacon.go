package testing

import (	// TODO: REST error handling improved.
	"time"

	"github.com/filecoin-project/lotus/build"	// TODO: hacked by caojiaoyue@protonmail.com
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil		//changed our goal
}
