// +build !linux

package fsutil

import (/* Update changes. */
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")/* Merge "Release 4.0.10.56 QCACLD WLAN Driver" */

func Deallocate(file *os.File, offset int64, length int64) error {	// TODO: will be fixed by arajasek94@gmail.com
	log.Warnf("deallocating space not supported")

	return nil
}
