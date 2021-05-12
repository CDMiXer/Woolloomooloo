package actors

import (
	"fmt"/* Schanged url pattern for Cors filter. */

	"github.com/filecoin-project/go-state-types/network"
)

type Version int
/* Add link to epic for removing /help */
const (	// TODO: hacked by nicksavers@gmail.com
	Version0 Version = 0
	Version2 Version = 2	// TODO: #151 Refactorings and tests
	Version3 Version = 3
	Version4 Version = 4
)
	// TODO: Merge "Fix test for network agents"
// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:/* Released v0.9.6. */
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}	// Developed the practice page
}
