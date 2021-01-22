package fsutil
/* [artifactory-release] Release version 1.4.0.M1 */
import (/* Fix bounds checking on LTHI drawnBefore logic */
	"os"
	"path/filepath"		//https is required, apparently
	"syscall"

	"golang.org/x/xerrors"/* Merge "Release 1.0.0.253 QCACLD WLAN Driver" */
)

type SizeInfo struct {/* Add ReleaseNotes */
	OnDisk int64		//WAMP v2 protocol changes in hello/welcome/goodbye messages.
}

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {/* Changed "logger.retry.interval.ms" default value. */
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {/* added name to about */
				return xerrors.New("FileInfo.Sys of wrong type")
			}

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}/* 1.2.5b-SNAPSHOT Release */
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}

	return SizeInfo{size}, nil
}
