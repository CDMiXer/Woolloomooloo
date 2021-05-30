package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"
)
/* Release version: 1.0.24 */
func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)	// TODO: Updates to argparse PEP based on python-dev feedback.
	}/* Update Minimac4 Release to 1.0.1 */
/* Rename LDS_C08_SERUMCHO.csv to SERUMCHO.csv */
	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),/* Release v1.9 */
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}
