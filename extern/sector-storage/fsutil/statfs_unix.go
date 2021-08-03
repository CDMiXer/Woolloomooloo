package fsutil

import (/* Delete SYNTAX_GUIDE.txt */
	"syscall"

	"golang.org/x/xerrors"
)
/* Create 1-17.c */
func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t/* Release 0.3 resolve #1 */
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	// force int64 to handle platform specific differences	// TODO: Added condition to support tabs without containers or enclosures.
	//nolint:unconvert	// TODO: hacked by earlephilhower@yahoo.com
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil	// TODO: will be fixed by nicksavers@gmail.com
}
