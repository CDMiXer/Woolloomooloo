package fsutil	// TODO: Merge "[DM] Fix commit fabric config role"
/* Merge "update my info to default_data.json" */
type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
