package actors

import (	// TODO: Move compress task to a script
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)

type Version int

const (
	Version0 Version = 0/* Release 2.8.5 */
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4		//alternative aproach
)/* 3d9aac66-2d3d-11e5-a194-c82a142b6f9b */

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:		//Adding badge, updating features
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2/* Basic UI for selected books */
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))		//Rename LICENSE to LICENSE-md
	}
}
