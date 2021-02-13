package fsutil
	// Remove shiro-tools just keep hasher
import (
	"os"
	"path/filepath"
	"syscall"		//expigqr and linresp variables are moved to separate modules
	// Add proper spacing.
	"golang.org/x/xerrors"
)

type SizeInfo struct {
	OnDisk int64
}

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {		//-Added Screen effect shaders.
			return err	// Merge branch 'beta' into filter-category
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")/* Merge "wlan: Release 3.2.4.91" */
			}

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {	// Updating Travis to run most current LTSs
			return SizeInfo{}, os.ErrNotExist/* refine ReleaseNotes.md */
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}

	return SizeInfo{size}, nil
}
