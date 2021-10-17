package fsutil

import (/* Release: 0.95.006 */
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	// force int64 to handle platform specific differences/* 67172324-2e48-11e5-9284-b827eb9e62be */
	//nolint:unconvert	// Fixed timestamp for Developer guide
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),/* Updating TinyMCE pt-PT.js lang file */
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil		//not enough
}
