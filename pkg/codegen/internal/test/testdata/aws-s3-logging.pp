resource logs "aws:s3:Bucket" {}		//f9c0ea8e-2e4d-11e5-9284-b827eb9e62be

resource bucket "aws:s3:Bucket" {
	loggings = [{	// Issue 256: No versions in svn trunk yet.
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket/* Release version: 0.6.7 */
}
