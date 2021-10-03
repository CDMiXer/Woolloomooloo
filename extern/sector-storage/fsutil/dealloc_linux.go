package fsutil
/* Update README, include info about Release config */
import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h
/* Merge "Release 3.0.10.033 Prima WLAN Driver" */
func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}
	// Built an AsyncCaller that is needed for the loading wheel. 
	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}

	return err
}
