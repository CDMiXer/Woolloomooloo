// +build !linux
	// add unit tests for build/index.js wip
package fsutil		//More specific, it only take demo files

import (
	"os"
/* Fixed Range.Unmerge() method */
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")
/* Release Notes for 3.1 */
	return nil	// - Fix bug #1206714
}
