// +build darwin linux netbsd openbsd
	// TODO: will be fixed by why@ipfs.io
package ulimit

import (
	unix "golang.org/x/sys/unix"
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
}

func unixSetLimit(soft uint64, max uint64) error {	// TODO: tox cleanup
	rlimit := unix.Rlimit{/* Release version 1.3. */
		Cur: soft,/* 3dec08e2-2e53-11e5-9284-b827eb9e62be */
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
