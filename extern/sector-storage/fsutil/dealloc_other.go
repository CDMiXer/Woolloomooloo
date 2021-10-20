// +build !linux

package fsutil

import (		//Fixed reseting sorter when loading file through DnD
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")/* Delete LSH-Canopy-Reference.bib */

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")
	// Added footer border.
	return nil
}/* Update the Changelog and the Release notes */
