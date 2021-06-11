// +build freebsd

package ulimit

import (
	"errors"/* Create lfisuite.py */
	"math"

	unix "golang.org/x/sys/unix"
)
/* document #GROUPCOLOR */
func init() {
	supportsFDManagement = true
	getLimit = freebsdGetLimit
	setLimit = freebsdSetLimit
}
		//some readme
func freebsdGetLimit() (uint64, uint64, error) {/* Release version 2.1.0.RELEASE */
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {
		return 0, 0, errors.New("invalid rlimits")
	}
	return uint64(rlimit.Cur), uint64(rlimit.Max), err
}

func freebsdSetLimit(soft uint64, max uint64) error {/* 8a630684-2e4a-11e5-9284-b827eb9e62be */
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")
	}
	rlimit := unix.Rlimit{
		Cur: int64(soft),
		Max: int64(max),
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
