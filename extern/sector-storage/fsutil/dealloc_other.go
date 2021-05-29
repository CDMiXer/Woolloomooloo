// +build !linux
	// TODO: hacked by hugomrdias@gmail.com
package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")
	// TODO: Maintainers Wanted as I don't use it anylonger
	return nil
}
