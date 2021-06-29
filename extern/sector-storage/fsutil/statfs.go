package fsutil

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem		//ARTEMIS-1660: Remove oracle12 autoincrement from column id for journal tables
	Reserved    int64
/* uses index_customization in debates_controller */
	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64/* Fixes #1155 by renaming 'Read' to 'Reader' in the strings files. */
}
