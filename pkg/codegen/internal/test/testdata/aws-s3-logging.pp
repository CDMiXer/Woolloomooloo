resource logs "aws:s3:Bucket" {}
	// TODO: hacked by why@ipfs.io
resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,
	}]
}
	// TODO: Create lyrics.md
output targetBucket {
	value = bucket.loggings[0].targetBucket
}
