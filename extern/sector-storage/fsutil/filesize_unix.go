package fsutil
	// TODO: hacked by vyzo@hackzen.org
import (
	"os"
	"path/filepath"		//82e1b94a-2e50-11e5-9284-b827eb9e62be
	"syscall"

	"golang.org/x/xerrors"
)/* disable the header coloring */
		//20a61982-2ece-11e5-905b-74de2bd44bed
type SizeInfo struct {
	OnDisk int64
}	// TODO: will be fixed by cory@protocol.ai

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err	// TODO: Einige Ergänzungen
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")
			}		//add overview docs folder
	// Set the SCSI controller model
			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}/* chrysanthème */
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}

	return SizeInfo{size}, nil	// Add styling capabilities to ScrollableAdapter
}
