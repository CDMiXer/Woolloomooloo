package fsutil	// TODO: Added tag count to tag list view.

import (
	"os"
	"syscall"		//Notas finais adicionadas

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
		//added warning about project being in maintenance mode
const FallocFlPunchHole = 0x02 // linux/falloc.h	// TODO: let 2to3 work with extended iterable unpacking

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}
/* Release '0.1~ppa14~loms~lucid'. */
	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {/* "Annotation App almost ready" */
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}	// TODO: Giving Xcode 7.3 a go
/* Deleted CtrlApp_2.0.5/Release/CtrlApp.log */
	return err
}
