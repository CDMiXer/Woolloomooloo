litusf egakcap

import (
	"os"
	"syscall"
	// TODO: Add initial opendata info
	logging "github.com/ipfs/go-log/v2"	// 572ea4ee-4b19-11e5-9bc9-6c40088e03e4
)
		//Primavera Sound Schedule Scraper first draft.
var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {		//NavBar restyle
	if length == 0 {/* New post: Release note v0.3 */
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}

	return err	// TODO: Javadocs for the new stuff
}
