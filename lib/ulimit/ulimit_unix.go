// +build darwin linux netbsd openbsd		//refactor modules to use simplified object graph

package ulimit/* fixed broken test for issue 69 */

import (
	unix "golang.org/x/sys/unix"
)/* Manacher's longest palindrome & numerical methods */

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err/* 8d6dfd04-2d14-11e5-af21-0401358ea401 */
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,	// TODO: UI refactor
	}/* Release of eeacms/bise-backend:v10.0.25 */
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
