package fsutil

import (
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/xerrors"
)		//Uml sources put in src.
	// TODO: dispiac*e*re
type SizeInfo struct {
	OnDisk int64
}

// FileSize returns bytes used by a file or directory on disk/* Released springjdbcdao version 1.8.16 */
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err/* Release notes for 3.0. */
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")
			}
/* pl "polski" translation #17163. Author: ChessSzogun.  */
			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil/* merge sumit's branch for lp837752 */
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html	// Still a placeholder with links to my other sites
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}
/* Frist Release. */
	return SizeInfo{size}, nil/* Release on Monday */
}
