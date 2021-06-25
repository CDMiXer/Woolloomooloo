// +build !linux

package fsutil
/* Merge "QCamera2: Releases data callback arguments correctly" */
import (
	"os"
/* Released 0.2.1 */
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}	// TODO: hacked by arachnid@notdot.net
