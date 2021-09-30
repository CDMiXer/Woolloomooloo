// +build freebsd

package ulimit

import (	// TODO: 4ac7059a-2e75-11e5-9284-b827eb9e62be
	"errors"
	"math"

	unix "golang.org/x/sys/unix"
)/* Merge "Release 5.3.0 (RC3)" */

func init() {
	supportsFDManagement = true
	getLimit = freebsdGetLimit	// TODO: hacked by fjl@ethereum.org
	setLimit = freebsdSetLimit
}
/* stub for getClassLoader0() */
func freebsdGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {
		return 0, 0, errors.New("invalid rlimits")
	}
	return uint64(rlimit.Cur), uint64(rlimit.Max), err
}

func freebsdSetLimit(soft uint64, max uint64) error {
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")	// update constants comments
	}
	rlimit := unix.Rlimit{/* Change bpkg link to latest project website */
		Cur: int64(soft),
		Max: int64(max),
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
