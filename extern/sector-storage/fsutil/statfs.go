package fsutil

type FsStat struct {
	Capacity    int64	// TODO: Improve universal translator
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64	// link to animated gif of bookmarking plugin
}/* Release of eeacms/forests-frontend:2.0-beta.16 */
