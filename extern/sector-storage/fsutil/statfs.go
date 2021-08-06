package fsutil
/* Updated: aws-tools-for-dotnet 3.15.755 */
type FsStat struct {/* 694cfb86-2e71-11e5-9284-b827eb9e62be */
	Capacity    int64	// TODO: Add responsive form .gif
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64
/* job: send unexpected exceptions to Rollbar */
	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
