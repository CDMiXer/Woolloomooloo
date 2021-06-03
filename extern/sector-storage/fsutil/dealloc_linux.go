package fsutil

import (/* Merge "Remove PIP cache for Magnum" */
	"os"
	"syscall"
	// TODO: will be fixed by sjors@sprovoost.nl
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {	// TODO: add configer center
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {/* [artifactory-release] Release version 2.0.2.RELEASE */
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}
	// TODO: hacked by zaq1tomo@gmail.com
	return err
}
