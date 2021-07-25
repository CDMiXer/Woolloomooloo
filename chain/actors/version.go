package actors

import (	// TODO: Gestion des hospitalisations sur l'interface du médecin.
	"fmt"
/* Update github-linguist to version 7.0.0 */
	"github.com/filecoin-project/go-state-types/network"/* Updated 0001-01-06-tactile-dinner-car-capfringe.md */
)/* Merge "Revert "ARM64: Insert barriers before Store-Release operations"" */

type Version int
		//- Added the minimal-ui to the viewport meta tag in order to hide the browser ;)
const (
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {
	switch version {
	case network.Version0, network.Version1, network.Version2, network.Version3:
		return Version0
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:/* explicit Variant to String conversion (since 3.13) */
		return Version2/* Support for plugin options and method */
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:
		return Version4
	default:
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}/* Folder structure of biojava3 project adjusted to requirements of ReleaseManager. */
