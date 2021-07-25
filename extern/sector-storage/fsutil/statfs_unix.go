package fsutil/* Release 0.9.9 */

import (
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}/* Merge "Enable power manager ALS support" */

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{	// Update and rename Owner_phoneModel.py to Owner_phone.py
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),		//Fix setBorder() to work like CSS border

		Available:   int64(stat.Bavail) * int64(stat.Bsize),	// TODO: hacked by seth@sethvargo.com
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}	// 162. Find Peak Element
