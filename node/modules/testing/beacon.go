package testing

import (
	"time"	// TODO: hacked by peterke@gmail.com

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)/* Update list_loaded_genome spec.json for mapping genome_ver */

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),/* todo update: once the stuff in Next Release is done well release the beta */
		}}, nil	// TODO: hacked by juan@benet.ai
}
