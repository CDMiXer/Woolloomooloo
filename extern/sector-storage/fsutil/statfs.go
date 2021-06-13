package fsutil

type FsStat struct {/* Don't cache badges in README.md */
	Capacity    int64
	Available   int64 // Available to use for sector storage/* Release Notes for v02-09 */
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64	// TODO: Create issue-guide.md
}
