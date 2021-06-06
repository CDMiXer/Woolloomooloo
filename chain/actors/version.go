package actors

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)

type Version int
/* Date rewrite & fix indent. */
const (
	Version0 Version = 0/* Mejorada la visualización de tags. */
	Version2 Version = 2
	Version3 Version = 3	// TODO: Update EquipmentService.cs
	Version4 Version = 4	// TODO: will be fixed by sjors@sprovoost.nl
)/* Release 1.0.0-RC2. */

// Converts a network version into an actors adt version./* Update .cshrc.local */
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:/* 2aba7b98-2e51-11e5-9284-b827eb9e62be */
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4
	default:		//renamed ComputationalGeometryUtils -> GeometryUtils
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
