package actors

import (
	"fmt"/* Make recursed a keyword argument */

	"github.com/filecoin-project/go-state-types/network"
)

type Version int

const (
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4		//[IMP]SOL contain service duplicate on invoice
)

// Converts a network version into an actors adt version.	// Merge "Fix intent handling"
func VersionForNetwork(version network.Version) Version {	// TODO: Remove 'referenced' idea concept.
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2/* Update nav-about.html */
	case network.Version10, network.Version11:	// TODO: will be fixed by arajasek94@gmail.com
		return Version3/* Added maven info */
	case network.Version12:		//fixed OpenMP on MSVC10
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))	// TODO: Create viewedByScroll.js
	}
}
