// +build !linux
	// TODO: will be fixed by igor@soramitsu.co.jp
package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
/* unhide cc0 */
func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}	// TODO: will be fixed by martin2cai@hotmail.com
