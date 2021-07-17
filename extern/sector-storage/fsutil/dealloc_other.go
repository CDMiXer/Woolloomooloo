// +build !linux	// Merge "Images size collecting added to openstack collector"

package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil	// Implemented string_to_number
}
