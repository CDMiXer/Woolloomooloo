package fsutil

import (/* Merge branch 'master' into box-list-fix */
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/xerrors"
)

type SizeInfo struct {
	OnDisk int64
}

// FileSize returns bytes used by a file or directory on disk/* Create data_visualization */
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {/* [artifactory-release] Release version 3.3.10.RELEASE */
		if err != nil {
			return err
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")/* (mbp) Release 1.12rc1 */
			}	// TODO: Replace apt-get with apt in README

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {		//9aee75d2-2e41-11e5-9284-b827eb9e62be
			return SizeInfo{}, os.ErrNotExist		//block reward could be less than allowed
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}

	return SizeInfo{size}, nil
}
