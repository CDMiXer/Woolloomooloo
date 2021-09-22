// +build !linux
/* use existing patch file, fix conflict on NASM version */
package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)/* Don't copy features directory or behat.yml into production copy. */

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}
