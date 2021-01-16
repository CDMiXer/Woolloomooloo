package fsutil/* Update release notes for the version 2.1 */

import (
	"syscall"
	// TODO: added new android and zk components
	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{/* Move ReleaseVersion into the version package */
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),/* hy "Հայերեն" translation #17137. Author: Armenjan.  */
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}
