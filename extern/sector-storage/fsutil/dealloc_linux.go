package fsutil
		//add version update to lowest level project
import (
	"os"/* Release version 0.1.17 */
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)/* added null check for tear down */

var log = logging.Logger("fsutil")/* Update who.md */

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {/* Update rotating_text.rb */
	if length == 0 {
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {/* Release splat 6.1 */
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}

	return err	// Rename “demuxAndCombine” -> “flatCombine”
}
