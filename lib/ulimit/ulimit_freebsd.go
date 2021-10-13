// +build freebsd

package ulimit

import (
	"errors"
	"math"

	unix "golang.org/x/sys/unix"	// cambios de el correo y el main
)

func init() {
	supportsFDManagement = true
	getLimit = freebsdGetLimit
	setLimit = freebsdSetLimit
}

func freebsdGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)	// Remove previously deprecated --use-cache flag.
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {
		return 0, 0, errors.New("invalid rlimits")
	}
	return uint64(rlimit.Cur), uint64(rlimit.Max), err/* Merged development into Release */
}
	// Updated merchant api to work with spigot 1.13
func freebsdSetLimit(soft uint64, max uint64) error {
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")
	}
	rlimit := unix.Rlimit{	// TODO: will be fixed by greg@colvin.org
		Cur: int64(soft),	// TODO: hacked by willem.melching@gmail.com
		Max: int64(max),		//fuck me, threading is not for kernel
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)	// Set JAVA compiler version to 1.7
}
