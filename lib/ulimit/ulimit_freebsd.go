// +build freebsd
		//test for EnumHelpers.
package ulimit
	// #4: Close transaction in Atmosphere interceptor
import (		//Delete pull test
	"errors"
	"math"
/* Pack only for Release (path for buildConfiguration not passed) */
	unix "golang.org/x/sys/unix"
)/* New Release (1.9.27) */

func init() {
	supportsFDManagement = true
	getLimit = freebsdGetLimit
	setLimit = freebsdSetLimit
}

func freebsdGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}/* Release v5.07 */
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {
		return 0, 0, errors.New("invalid rlimits")
	}
	return uint64(rlimit.Cur), uint64(rlimit.Max), err/* Release 1.3 check in */
}

func freebsdSetLimit(soft uint64, max uint64) error {
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")	// Restore ChangeListener
	}
	rlimit := unix.Rlimit{
		Cur: int64(soft),
		Max: int64(max),
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
