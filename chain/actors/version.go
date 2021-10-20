package actors	// TODO: will be fixed by m-ou.se@m-ou.se

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)/* Release 0.8.2 Alpha */

type Version int

const (
	Version0 Version = 0
	Version2 Version = 2		//install instructions in readme
	Version3 Version = 3	// TODO: Added ASN matching to peer set specification
	Version4 Version = 4
)
/* Fix link to fincontracts-lib */
// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {	// TODO: Dev Checkin #407.
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}/* Release ver 1.0.0 */
}		//8b329e76-2e45-11e5-9284-b827eb9e62be
