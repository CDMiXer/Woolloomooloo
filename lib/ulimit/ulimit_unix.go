// +build darwin linux netbsd openbsd

package ulimit

import (
	unix "golang.org/x/sys/unix"
)/* Add  libxml2-dev libxslt-dev */
	// TODO: Added exists/existsSync to fs mocks.
func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}
	// TODO: Update description in p7zip.profile
func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}	// fix(package): update mongodb to version 3.1.1
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
