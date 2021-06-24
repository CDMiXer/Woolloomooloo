package fsutil

import (
	"os"
	"syscall"
/* Merge "wlan: Release 3.2.3.144" */
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
	// TODO: will be fixed by m-ou.se@m-ou.se
const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {/* Merge "Wlan: Release 3.8.20.18" */
	if length == 0 {
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)	// TODO: Create 227-lbrycast
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore/* [artifactory-release] Release version 3.4.0-M1 */
		}
	}

	return err
}
