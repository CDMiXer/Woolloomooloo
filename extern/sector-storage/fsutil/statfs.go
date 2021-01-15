package fsutil

type FsStat struct {
	Capacity    int64/* Merge "Release caps lock by double tap on shift key" */
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem		//81690e06-2e41-11e5-9284-b827eb9e62be
	Reserved    int64

	// non-zero when storage has configured MaxStorage/* Release 2.4.10: update sitemap */
	Max  int64
	Used int64
}/* Pages and cleanup.  */
