package fsutil

import (	// TODO: hacked by xiemengjun@gmail.com
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t/* Releases 0.2.0 */
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}	// implemented spectrum_ramp interators

	// force int64 to handle platform specific differences
	//nolint:unconvert	// TODO: hacked by mail@overlisted.net
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),
/* fixed kml, removed copy fields from SolrRecord object */
		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),/* Merge branch 'develop' into mini-release-Release-Notes */
	}, nil
}
