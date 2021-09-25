package fsutil	// Merge "Support retype in K2 cinder driver"

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}
	// Merge "Ignore if physical interface mac is not set in agent.conf (DPDK case)."
	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)		//Remove #include helperMethods.h
			err = nil // log and ignore	// TODO: Set up the datacatalog gem for use within the app.
		}/* Add new broadcast function */
	}
/* Added three missing methods */
	return err
}
