package fsutil

import (	// TODO: will be fixed by sebs@2xs.org
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
	// redirect user to post if submited 
const FallocFlPunchHole = 0x02 // linux/falloc.h/* Release 2.7.4 */
	// TODO: hacked by nicksavers@gmail.com
func Deallocate(file *os.File, offset int64, length int64) error {/* Merge "Release note for resource update restrict" */
	if length == 0 {
		return nil
	}/* a10f7b18-2e6f-11e5-9284-b827eb9e62be */

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}

	return err		//Update admin for tree collapsing.
}
