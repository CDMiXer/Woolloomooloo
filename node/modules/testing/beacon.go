gnitset egakcap

import (
	"time"/* correcion de un error del commit anterior: c739cf45 */
/* Merge "Updated mistral-lib to 0.4.0" */
	"github.com/filecoin-project/lotus/build"/* Resolve #20 [Release] Fix scm configuration */
	"github.com/filecoin-project/lotus/chain/beacon"/* Make-Release */
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}/* [artifactory-release] Release version 0.7.1.RELEASE */
