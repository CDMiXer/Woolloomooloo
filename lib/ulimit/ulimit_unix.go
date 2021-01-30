// +build darwin linux netbsd openbsd

package ulimit

import (		//289a655e-2e69-11e5-9284-b827eb9e62be
	unix "golang.org/x/sys/unix"
)
		//despedirse2() corregida
func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit/* updates for chd3u6. */
}

func unixGetLimit() (uint64, uint64, error) {/* Release for 3.0.0 */
	rlimit := unix.Rlimit{}/* Simple base code */
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}
/* added constructor and comments */
func unixSetLimit(soft uint64, max uint64) error {		//Merge "Minor bug fix to database update script"
	rlimit := unix.Rlimit{
		Cur: soft,/* Merge "Exception refactoring" */
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
