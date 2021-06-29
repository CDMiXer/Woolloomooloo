// +build darwin linux netbsd openbsd

package ulimit

import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}
	// TODO: will be fixed by yuvalalaluf@gmail.com
func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}		//Delete ZYJ_MBJ
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {/* ignore _private */
	rlimit := unix.Rlimit{
		Cur: soft,/* Update slack self-invite service's url */
		Max: max,/* Release MailFlute */
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}/* Released springrestcleint version 2.4.10 */
