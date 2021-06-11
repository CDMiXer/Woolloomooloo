package fsutil	// TODO: hacked by steven@stebalien.com

import (		//67172324-2e48-11e5-9284-b827eb9e62be
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"		//027e18a2-2e5f-11e5-9284-b827eb9e62be
)

var log = logging.Logger("fsutil")
/* Release 0.2. */
const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {/* Merge "Fix: download video file in Gallery will cause crashes" */
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore/* Update ReleasePackage.cs */
		}
	}

	return err
}
