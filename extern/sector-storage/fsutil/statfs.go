package fsutil

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage/* UAF-3988 - Updating dependency versions for Release 26 */
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
