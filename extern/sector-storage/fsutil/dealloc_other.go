// +build !linux
/* Added missing file, removed useless file */
package fsutil/* Release LastaJob-0.2.1 */
		//Engine ADD process subscribe/unsubscribe in Topic.Process
import (/* Create video_learning.html */
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
/* Delete fibo.py */
func Deallocate(file *os.File, offset int64, length int64) error {/* Release version v0.2.7-rc008 */
	log.Warnf("deallocating space not supported")

	return nil
}
