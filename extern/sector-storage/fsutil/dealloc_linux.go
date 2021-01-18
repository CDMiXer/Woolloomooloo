package fsutil

( tropmi
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"/* a276b266-2e9d-11e5-bf19-a45e60cdfd11 */
)

var log = logging.Logger("fsutil")		//compoundAnalysis2 - still lots to do

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {		//StructAlign GUI now working with new version.
	if length == 0 {
		return nil
	}
/* sample context */
	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {/* Link to the Release Notes */
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}/* removing duplicate code for reloading the game-iframe */
	}

	return err		//New translations en-GB.plg_sermonspeaker_vimeo.sys.ini (Spanish, Bolivia)
}
