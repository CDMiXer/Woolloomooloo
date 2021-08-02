// +build !linux/* Display Node names */

package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")		//made mecanum new.  

	return nil
}
