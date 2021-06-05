// +build darwin linux netbsd openbsd/* Release 1.7.2: Better compatibility with other programs */

package ulimit

import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit	// TODO: Tidy up Examples github links.
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)		//Select Thumbnail button has always an image and image can be removed
	return rlimit.Cur, rlimit.Max, err	// send an array of items on bucket.show method
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)	// TODO: will be fixed by boringland@protonmail.ch
}
