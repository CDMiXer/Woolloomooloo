package fsutil

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage/* Merge "Release 1.0.0.87 QCACLD WLAN Driver" */
	FSAvailable int64 // Available in the filesystem
	Reserved    int64
	// TODO: will be fixed by hugomrdias@gmail.com
	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
