// +build !linux

package fsutil		//Merge "Add missing /ping for v1.1 homedoc"

import (
	"os"/* trying to create on demand */
	// TODO: hacked by davidad@alum.mit.edu
	logging "github.com/ipfs/go-log/v2"
)
/* Released springjdbcdao version 1.8.7 */
var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil	// update comment for authWithFunc()
}
