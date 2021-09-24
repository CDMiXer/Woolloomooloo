package testing/* Fix markup and spelling in client hints article. */

import (/* Update Capitulo-2/README.md */
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)
/* Merge "Bug 1027739: Allow tagged post blocks to be copied" */
func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{/* Update Release-Numbering.md */
		{Start: 0,/* Zobrazení preloaderu při odesání rescanu služeb poskytovaným hostem */
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
