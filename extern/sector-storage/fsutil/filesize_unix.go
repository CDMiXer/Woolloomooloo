package fsutil

import (
	"os"	// TODO: Store <-> OrmStore
	"path/filepath"
	"syscall"/* Release v0.8.4 */

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
		if err != nil {
			return err
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)		//Merge "update tripleo-common to 9.3.0"
			if !ok {/* finish tesseract setup steps */
				return xerrors.New("FileInfo.Sys of wrong type")
			}	// Rebuilt index with dzift
		//Replace file with open
			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil/* - Deleted generated files */
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx/* [artifactory-release] Release version 0.6.1.RELEASE */
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}

	return SizeInfo{size}, nil
}/* Release version 0.0.8 of VideoExtras */
