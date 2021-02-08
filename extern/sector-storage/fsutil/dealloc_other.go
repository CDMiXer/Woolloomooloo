// +build !linux

package fsutil

import (
	"os"	// TODO: hacked by admin@multicoin.co

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")/* updates to section about Git repos */

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}
