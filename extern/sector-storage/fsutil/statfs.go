package fsutil	// TODO: hacked by cory@protocol.ai

type FsStat struct {		//Delete royal_preloader.min.css
	Capacity    int64/* Update Orchard-1-9-Release-Notes.markdown */
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage/* Release of primecount-0.16 */
	Max  int64
	Used int64
}
