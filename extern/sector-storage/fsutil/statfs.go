package fsutil

type FsStat struct {
	Capacity    int64	// TODO: Merge lp:~gundlach/nova/missing_ec2_url_endpoints
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem	// TODO: Update and rename 404._ to 404.html
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
