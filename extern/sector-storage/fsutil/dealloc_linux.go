package fsutil	// Create Installer.php
	// Tagged by Jenkins Task SVNTagging. Build:jenkins-YAKINDU_SCT2_CI-1118.
import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"/* points to real documentation */
)

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h/* 7912092c-2e43-11e5-9284-b827eb9e62be */

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil/* OpenTK svn Release */
	}/* New classes copied from JCommon. */

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}
/* Fixes naming of config properties */
	return err/* Issue #4512 closeout: Make ZipImport.get_filename() a public method */
}
