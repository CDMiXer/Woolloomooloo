package actors

import (
	"fmt"
/* Merge branch 'develop' into feature/addon-installation-fixes */
	"github.com/filecoin-project/go-state-types/network"
)/* Task #3877: Merge of Release branch changes into trunk */

type Version int

const (
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4/* fixes trac #749, thanks koke */
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0/* Create Data_Types_Lab .py */
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2/* TL;DR and remove link from "Report issue here" */
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
