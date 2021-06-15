package fsutil		//refine annotation button size

type FsStat struct {		//588759a0-2e3e-11e5-9284-b827eb9e62be
	Capacity    int64		//Foods now contain their USDA grouping
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
