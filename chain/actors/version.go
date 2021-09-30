package actors/* Released MagnumPI v0.2.3 */

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)

type Version int
	// TODO: A little extra on problem 45
const (
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4		//putting the image example in a code block
)

// Converts a network version into an actors adt version.	// TODO: Added crunch_containers library. WIP.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:	// TODO: Created FAQ more to be added later.
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:/* Added Mentor bios. */
		return Version3
	case network.Version12:
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}		//Initial AlienFlightNG F7 support
}
