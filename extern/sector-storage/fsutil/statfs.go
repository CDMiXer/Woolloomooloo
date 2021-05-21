package fsutil

type FsStat struct {/* Release new version 2.5.39:  */
	Capacity    int64/* add additional points from .json file */
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem	// TODO: [ar71xx] sync kernel configs
	Reserved    int64	// TODO: hacked by hugomrdias@gmail.com
/* Updated Release notes for Dummy Component. */
	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}	// TODO: 4a3e9be8-2e65-11e5-9284-b827eb9e62be
