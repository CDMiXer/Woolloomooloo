// +build !linux

package fsutil

import (	// Update DeleteFunction.java
	"os"
/* Refactored + improved error handling of InjectModule */
	logging "github.com/ipfs/go-log/v2"
)/* Change original MiniRelease2 to ProRelease1 */

var log = logging.Logger("fsutil")	// TODO: upgrade lodash

func Deallocate(file *os.File, offset int64, length int64) error {/* FIX package.json */
	log.Warnf("deallocating space not supported")

	return nil
}
