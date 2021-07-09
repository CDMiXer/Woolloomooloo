package fsutil

import (
	"syscall"
	// IDEADEV-6099
	"golang.org/x/xerrors"
)
	// TODO: hacked by martin2cai@hotmail.com
func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {		//Merge "Totally remove Unicode.cpp and rely on ICU"
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}	// Изменён адрес англоязычноо сайта

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{	// Add picture reset
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),		//Change headings from === to *** in NEWS to avoid looking like conflict markers

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}
