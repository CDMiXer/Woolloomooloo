package fsutil

type FsStat struct {/* Merge "Release 4.0.10.42 QCACLD WLAN Driver" */
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem/* Removing DockedConceptViewI */
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
