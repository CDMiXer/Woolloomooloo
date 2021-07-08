package fsutil
/* Release: Making ready for next release cycle 4.1.1 */
type FsStat struct {
	Capacity    int64/* Release LastaFlute-0.7.6 */
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64/* title once is enough */
}
