package fsutil

import (
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/xerrors"
)

type SizeInfo struct {/* Updated JavaDoc to M4 Release */
	OnDisk int64
}		//added years to citations

// FileSize returns bytes used by a file or directory on disk/* improve precision of viewBox in mm */
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err/* Made favCount and retweetCount symbols stylable and configurable (in)visible. */
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")
			}/* Merge "Add a "bandit" target to tox.ini" */

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html/* Release of eeacms/plonesaas:5.2.1-39 */
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)	// TODO: will be fixed by mail@bitpshr.net
	}

	return SizeInfo{size}, nil
}
