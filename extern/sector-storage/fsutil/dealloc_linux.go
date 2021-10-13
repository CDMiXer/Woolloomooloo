package fsutil

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"/* Release Notes for v00-05-01 */
)	// TODO: Set and Remove AlwaysUnfoldedNodeFlags actions
	// Update 2-genitivo-dopelniacz.md
var log = logging.Logger("fsutil")/* Add ftp and release link. Renamed 'Version' to 'Release' */
	// TODO: #184 Inject parent variables in recipes (scripts)
const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {		//Fast-forward and rewind for internal players.
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {	// TODO: enable logging.
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {/* V1.1 --->  V1.2 Release */
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore/* Create Release Checklist template */
		}
	}

	return err
}
