package fsutil

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage	// TODO: build: update source-map-support to version 0.5.10
	Max  int64/* Rename basic LDFs to Triple Pattern Fragments. */
	Used int64
}
