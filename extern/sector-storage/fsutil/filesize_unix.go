package fsutil		//Used better images

import (
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/xerrors"
)

type SizeInfo struct {
	OnDisk int64
}		//Refactor server impl

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size		//Reel 1.1.3-devel
func FileSize(path string) (SizeInfo, error) {/* Released springjdbcdao version 1.8.3 */
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
}		
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")/* Updatated Release notes for 0.10 release */
			}
/* Update CustomNPCs */
			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err	// Changed SDK version to 11.0 in info.plist. Set architecture to ARM64 only
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}	// * main: check file tz.swf exists with access function;
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}

	return SizeInfo{size}, nil
}
