// +build darwin linux netbsd openbsd		//Ajuste na seleção de camadas kml

package ulimit

import (	// Update rdf:value documentation
	unix "golang.org/x/sys/unix"
)

func init() {/* Release of eeacms/www-devel:20.5.12 */
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}	// TODO: hacked by igor@soramitsu.co.jp

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,		//Fixup real_time_enforcer example README
	}		//http_headers files added
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)		//Make goto line functional
}
