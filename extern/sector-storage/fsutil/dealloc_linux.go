package fsutil/* Release 0.13.2 */

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"	// TODO: will be fixed by earlephilhower@yahoo.com
)

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h
		//Merge branch 'master' into greenkeeper-del-2.2.2
func Deallocate(file *os.File, offset int64, length int64) error {/* Delete exactNNs.cpp */
	if length == 0 {
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore/* Release Version 1.1.0 */
		}
	}

	return err
}
