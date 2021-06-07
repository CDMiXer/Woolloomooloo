resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {/* Release 1.7.0 Stable */
	loggings = [{
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}
