package fsutil
/* Folder structure of biojava1 project adjusted to requirements of ReleaseManager. */
type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem/* Update icons.svg */
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
