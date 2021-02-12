// +build darwin linux netbsd openbsd

package ulimit	// TODO: will be fixed by vyzo@hackzen.org
/* Bump version to coincide with Release 5.1 */
import (
	unix "golang.org/x/sys/unix"		//Delete 03-config.png
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}	// TODO: Merge "Save and restore brightness on orientation changes."

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{/* Create datastructure-polymer.h */
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
