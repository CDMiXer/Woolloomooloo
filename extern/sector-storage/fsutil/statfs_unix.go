package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"	// TODO: merged again.
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),		//Merge "Add import test"

		Available:   int64(stat.Bavail) * int64(stat.Bsize),	// TODO: Camara de fotos con comprobaciones de memoria externa. 
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),/* Excplicit the tag limit #1815 related */
	}, nil		//Merge "[MIPS64] Temporary placeholder build, to allow other projects to build"
}
