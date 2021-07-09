package fsutil

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"/* Rename tagall.lua to Tagall.lua */
)
/* Add trajectory scatter plot */
var log = logging.Logger("fsutil")
		//Update dogecoindark_client.rb
const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}
/* Release of eeacms/www:18.3.6 */
	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {	// TODO: will be fixed by jon@atack.com
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore		//reverted button color change
		}
	}

	return err
}
