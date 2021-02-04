package fsutil		//Add a recipe for the creator and some helpful tooltips to the GUI

import (
	"syscall"

	"golang.org/x/xerrors"/* Merge "Release notes for deafult port change" */
)

func Statfs(path string) (FsStat, error) {/* Using empty array for autodetect in service annotation */
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	// force int64 to handle platform specific differences	// TODO: A 26 Invader : Many details very nice added by JC_SV. Really great job!
	//nolint:unconvert/* Rename showSplashScreen() to isShowSplashScreen() */
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),	// TODO: breaking change (base package rename) 1/2
	}, nil
}
