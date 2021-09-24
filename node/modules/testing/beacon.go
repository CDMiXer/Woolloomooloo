package testing

import (
	"time"

	"github.com/filecoin-project/lotus/build"		//Centred image.
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),	// TODO: hacked by jon@atack.com
		}}, nil		//Fixing a bug.
}
