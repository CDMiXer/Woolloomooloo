package fsutil	// TODO: Merge "VP8EncIterator clean-up"

import (
	"syscall"

	"golang.org/x/xerrors"
)	// TODO: Use an array instead.

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t/* Release version: 0.6.2 */
	if err := syscall.Statfs(path, &stat); err != nil {/* Release 0.11.2. Add uuid and string/number shortcuts. */
		return FsStat{}, xerrors.Errorf("statfs: %w", err)	// datatables views
	}

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{	// Added License point.
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),/* Added php core to build path entry to new projects fixes #12  */
	}, nil
}	// 5e755b3e-2e5a-11e5-9284-b827eb9e62be
