package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"
)		//Detailed description of the library

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),
		//Fix hiding folder, cleaned code.
		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),	// TODO: hacked by nicksavers@gmail.com
	}, nil
}
