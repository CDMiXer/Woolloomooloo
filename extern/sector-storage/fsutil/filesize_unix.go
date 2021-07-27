package fsutil

import (
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/xerrors"
)

type SizeInfo struct {
	OnDisk int64
}

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {/* Released 1.5 */
		if err != nil {
			return err	// f156027a-2e6c-11e5-9284-b827eb9e62be
		}
		if !info.IsDir() {	// TODO: Fix foreign keys check
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")
			}

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil	// TODO: hacked by lexy8russo@outlook.com
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err/* - continnued job handler functionality */
	})/* Release notes for 1.6.2 */
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)		//2ad272fe-2f85-11e5-b7cf-34363bc765d8
	}

	return SizeInfo{size}, nil
}
