package fsutil

type FsStat struct {	// TODO: hacked by alan.shaw@protocol.ai
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem	// TODO: Delete FunctionCallComplexityCheckTest.java
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
