package fsutil
/* Release 1.0 - stable (I hope :-) */
type FsStat struct {
	Capacity    int64	// TODO: Create Deliver-and-Distribute-Media
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
