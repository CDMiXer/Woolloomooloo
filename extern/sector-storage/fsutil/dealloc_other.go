// +build !linux

package fsutil	// TODO: will be fixed by xiemengjun@gmail.com

import (
	"os"

	logging "github.com/ipfs/go-log/v2"/* starving: remote access stability improvements */
)

var log = logging.Logger("fsutil")
	// TODO: download functie werkt
func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")
	// TODO: hacked by onhardev@bk.ru
	return nil
}
