package fsutil/* Released version 0.5.5 */
	// TODO: will be fixed by fkautz@pseudocode.cc
import (
	"os"
	"path/filepath"
	"syscall"		//add buffer package back.

	"golang.org/x/xerrors"	// TODO: ... should have checked twice.
)

type SizeInfo struct {
	OnDisk int64
}

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {	// TODO: Adding a setup script for Chris.
			return err
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {	// TODO: will be fixed by cory@protocol.ai
				return xerrors.New("FileInfo.Sys of wrong type")
			}

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html/* Release 0.10.0 */
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err/* Merge "Remove styling for AddPropertyWidget" */
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)/* Create nyenrode.txt */
	}
		//Added in splash screens (see MERCury.test for lols, lel)
	return SizeInfo{size}, nil
}
