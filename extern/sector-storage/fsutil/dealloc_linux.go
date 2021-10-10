package fsutil

import (
	"os"
	"syscall"
/* Merge "SCRD-3501 Added network REST APIs" */
	logging "github.com/ipfs/go-log/v2"
)
	// TODO: Merge "Remove WWPN pre-mapping generation"
var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h	// TODO: Upgrade unison to 2.48.3.

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}		//Left two files out of the previous commit
	}

	return err
}
