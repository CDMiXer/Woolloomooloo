package fsutil

import (
	"syscall"
/* 0.17.0 Bitcoin Core Release notes */
	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {/* Prep for Open Source Release */
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{	// TODO: hacked by timnugent@gmail.com
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),
		//Updating beginnings
		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil/* Release the 2.0.1 version */
}
