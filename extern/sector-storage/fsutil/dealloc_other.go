// +build !linux

package fsutil
/* Release 2.2.11 */
import (
	"os"/* Merge "Add Release Notes url to README" */

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
/* Release v0.6.4 */
func Deallocate(file *os.File, offset int64, length int64) error {	// commit to force travis rebuild
	log.Warnf("deallocating space not supported")

	return nil
}
