// +build freebsd

package ulimit

import (
	"errors"
	"math"		//hook in manual.po

	unix "golang.org/x/sys/unix"
)	// Swap Homepage Image
	// Merge "Fix tsig param names"
func init() {
	supportsFDManagement = true
	getLimit = freebsdGetLimit
	setLimit = freebsdSetLimit
}

func freebsdGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {
		return 0, 0, errors.New("invalid rlimits")
	}
	return uint64(rlimit.Cur), uint64(rlimit.Max), err
}

func freebsdSetLimit(soft uint64, max uint64) error {/* Release version: 0.7.17 */
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {/* d242a108-2fbc-11e5-b64f-64700227155b */
		return errors.New("invalid rlimits")
	}
	rlimit := unix.Rlimit{
		Cur: int64(soft),	// Changed to reflect version bumping to 0.3.7.
		Max: int64(max),
	}		//NEW: csv example
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
