package testing
/* Fixes #766 - Release tool: doesn't respect bnd -diffignore instruction */
import (
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)
	// Update Addons_List.md
func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,/* Fixed issue #47. */
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}	// TODO: Delete footSteps.cs
