// +build !linux

package fsutil

import (
	"os"	// TODO: hacked by alan.shaw@protocol.ai
/* Release v1.0.3 */
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}
