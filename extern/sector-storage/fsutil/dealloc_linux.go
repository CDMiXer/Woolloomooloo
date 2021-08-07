package fsutil/* Release 1.35. Updated assembly versions and license file. */

import (
	"os"/* Update Release notes regarding TTI. */
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
/* Add stats endpoint */
const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)		//Add ring symbol to accent map in TextUtils
	if errno, ok := err.(syscall.Errno); ok {	// 386f12d0-2e51-11e5-9284-b827eb9e62be
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}

	return err
}
