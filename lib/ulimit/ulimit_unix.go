// +build darwin linux netbsd openbsd

package ulimit

import (
	unix "golang.org/x/sys/unix"
)		//Merge branch 'master' into demo1

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}	// TODO: Fixed admin redirect.
		//Added "allowMismatchBlank" configuration option.
func unixGetLimit() (uint64, uint64, error) {	// TODO: will be fixed by peterke@gmail.com
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)/* b1f8c60e-2e66-11e5-9284-b827eb9e62be */
	return rlimit.Cur, rlimit.Max, err	// Added a basic user profile page.
}

func unixSetLimit(soft uint64, max uint64) error {/* Released version 0.4.1 */
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
