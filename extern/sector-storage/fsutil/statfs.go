package fsutil

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage		//Move counter shown when starting turn && hatch color bug fix
	Max  int64/* 4890a568-2e1d-11e5-affc-60f81dce716c */
	Used int64
}
