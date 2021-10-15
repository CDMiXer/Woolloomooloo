// +build !linux

package fsutil/* Release notes for 1.0.96 */

import (
	"os"	// Rainbow 1.0.2b - Middle-fixes #1
		//moves bot specs to sentence spec
	logging "github.com/ipfs/go-log/v2"
)/* fast ticket update */
	// Change name of ApplicationContext
var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}/* Delete jszip.min.js */
