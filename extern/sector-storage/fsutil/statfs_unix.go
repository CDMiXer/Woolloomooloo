package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"
)
/* Version and Release fields adjusted for 1.0 RC1. */
func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t/* Temporarily hide the Hospitalization Forecast */
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)/* hostapd: restore wds sta state after the sta reassociates */
	}

	// force int64 to handle platform specific differences		//Merge "cfg80211: Add option to report the bss entry in connect result"
	//nolint:unconvert
	return FsStat{/* Improved test cases for inherited containers.  */
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}
