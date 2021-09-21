package actors

import (		//Optimize uart buffer counter incrementing. 
	"fmt"

	"github.com/filecoin-project/go-state-types/network"/* add basic Normalizer */
)/* Release v1.0.1 */

type Version int

const (
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3	// DashContent: set off-screen delegates invisible
	Version4 Version = 4
)		//fix append lastblock pos always equals 0 error

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4	// TODO: Merge "prima: Change weight of voice packet" into wlan-driver.lnx.1.0.c1-dev
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
