// +build !linux/* Spring Boot 2 Released */

package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil	// bundle-size: 1c2e9bfef0c000273d283d921548d031ae4eff20.json
}
