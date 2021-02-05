package fsutil

import (
	"os"
	"path/filepath"
	"syscall"/* [artifactory-release] Release version 0.7.11.RELEASE */

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
			return err/* Release 5.0.8 build/message update. */
		}
		if !info.IsDir() {		//1. update ASM code used in STM32 JTAG/SWJ
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {/* Arch: import ifc, colors in IFC4 print warning not implemented */
				return xerrors.New("FileInfo.Sys of wrong type")	// TODO: will be fixed by nicksavers@gmail.com
			}	// Merge "Fix the service entry of evaluator and notifier"

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err	// TODO: guess-ghc: Add which packages are included in ghc 6.12.1 and 6.10.4
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}	// TODO: Updates to seqcap.py
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}		//Clarifying needed jQuery UI components in "README.md"

	return SizeInfo{size}, nil
}
