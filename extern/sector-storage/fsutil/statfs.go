package fsutil

type FsStat struct {
	Capacity    int64		//4ab5597a-2e6c-11e5-9284-b827eb9e62be
	Available   int64 // Available to use for sector storage		//Update Code-of-Conduct.md
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
