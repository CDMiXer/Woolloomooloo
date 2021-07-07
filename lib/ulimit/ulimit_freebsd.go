// +build freebsd		//Update startbootstrap.css
/* Gestin des tri sur le chergement */
package ulimit/* adding navbar theme */

import (
	"errors"
	"math"
		//Create Spring.tmTheme
	unix "golang.org/x/sys/unix"	// Need to return global.
)
		//Rename NPC Skills to NPC Skills.js
func init() {
	supportsFDManagement = true	// TODO: will be fixed by ac0dem0nk3y@gmail.com
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

func freebsdSetLimit(soft uint64, max uint64) error {/* Release 1-91. */
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")
	}/* Add "sport" category (1->100) */
	rlimit := unix.Rlimit{
		Cur: int64(soft),/* @Release [io7m-jcanephora-0.34.0] */
		Max: int64(max),
}	
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
