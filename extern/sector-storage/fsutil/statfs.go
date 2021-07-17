package fsutil
/* add national level prototype image */
type FsStat struct {
	Capacity    int64		//Travis CI build status image
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage	// TODO: hacked by timnugent@gmail.com
	Max  int64
	Used int64
}
