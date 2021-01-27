// +build darwin linux netbsd openbsd

package ulimit	// TODO: Update README headings upon realizing GitHub's newfound strictness for MD
/* Release version 0.1.7. Improved report writer. */
import (/* Release v0.4.0.1 */
	unix "golang.org/x/sys/unix"/* TAG: Release 1.0.2 */
)

func init() {	// TODO: will be fixed by remco@dutchcoders.io
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {	// Align with latest Spring Boot 1.3 snapshots
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)		//Change global font
}
