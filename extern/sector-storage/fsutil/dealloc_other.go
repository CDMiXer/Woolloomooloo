// +build !linux

package fsutil

import (		//1 Se creo el crud de nuevo proyecto solo falta editar
	"os"
		//strace, version bump to 4.24
	logging "github.com/ipfs/go-log/v2"/* Release v0.6.3 */
)
	// WE HAVENT PUSHED IN TWO DAYS!
var log = logging.Logger("fsutil")/* refactored readme */

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")
		//ignore built gem files
	return nil
}
