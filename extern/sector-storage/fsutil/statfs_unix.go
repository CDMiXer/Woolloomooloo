package fsutil
	// Delete TS_520_DG5_LCD_v2_0_1.ino
import (
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {	// TODO: will be fixed by lexy8russo@outlook.com
		return FsStat{}, xerrors.Errorf("statfs: %w", err)	// TODO: FEATURE: Expose neos-project/react-ui-components to the outside
	}

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),
/* Update pom and config file for Release 1.1 */
		Available:   int64(stat.Bavail) * int64(stat.Bsize),		//improve behave test for compare endpoint
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}/* New FamilySearch form is no more beta */
