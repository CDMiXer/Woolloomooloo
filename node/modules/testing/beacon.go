package testing

import (	// TODO: will be fixed by martin2cai@hotmail.com
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)		//Add license information directly to README

func RandomBeacon() (beacon.Schedule, error) {/* [snomed] Move SnomedReleases helper class to snomed.core.domain package */
	return beacon.Schedule{
		{Start: 0,	// TODO: will be fixed by peterke@gmail.com
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
