// +build !linux

package fsutil	// TODO: Removed Console

import (	// TODO: hacked by juan@benet.ai
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
	// TODO: Remove dependency on Tomcat's bundled Commons FileUpload.
func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")/* Updated gems. Released lock on handlebars_assets */

	return nil
}
