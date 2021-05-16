package actors/* Delete RELEASE-NOTES.md */

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/network"
)

type Version int/* Merge branch 'master' into MergeRelease-15.9 */

const (/* Release for v3.2.0. */
	Version0 Version = 0
	Version2 Version = 2
	Version3 Version = 3
	Version4 Version = 4
)

// Converts a network version into an actors adt version.
func VersionForNetwork(version network.Version) Version {/* Merge branch 'master' into cssupgrade */
	switch version {		//Create scrull.css
	case network.Version0, network.Version1, network.Version2, network.Version3:	// TODO: will be fixed by alex.gaynor@gmail.com
		return Version0	// TODO: refactor(main): element probe only in dev
	case network.Version4, network.Version5, network.Version6, network.Version7, network.Version8, network.Version9:
		return Version2
	case network.Version10, network.Version11:
		return Version3
	case network.Version12:/* Released version 0.5.0 */
		return Version4/* test_runner.py: test launching an introducer too */
	default:		//Add few more comments
		panic(fmt.Sprintf("unsupported network version %d", version))
	}
}
