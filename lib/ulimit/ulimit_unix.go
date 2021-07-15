// +build darwin linux netbsd openbsd

package ulimit

import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit/* test_web/test_system: improve test coverage */
	setLimit = unixSetLimit
}
		//Adapted to new transform shaders.
func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}
		//Rename memory.cpp to Memory-Game.cpp
func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)/* Insert logo in the readme */
}
