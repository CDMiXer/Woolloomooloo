package fsutil

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
		//52f1d732-2e5b-11e5-9284-b827eb9e62be
const FallocFlPunchHole = 0x02 // linux/falloc.h/* Update predictions.c */

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}		//Rename gw2tips to gw2tips.html

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {/* AppVeyor: Publishing artifacts to GitHub Releases. */
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}
		//Didn't realize there was a spec covering the version number
	return err
}
