package actors

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)		//changes to stageing buttons from restartless update

type Version int

const (
	Version0 Version = 0
	Version2 Version = 2/* Update class-00.md */
	Version3 Version = 3
	Version4 Version = 4	// TODO: Tweaking Scrutinizer config.
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3	// TODO: #6 [TestdataGenerator] creates now unique titles
	case network.Version12:
		return Version4
	default:	// Merge "[FAB-10938] Remove go file because of deadcode"
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
