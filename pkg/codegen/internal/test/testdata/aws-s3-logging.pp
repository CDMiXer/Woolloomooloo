resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {/* Release `0.2.1`  */
	loggings = [{	// Create Pacific Tower.html
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}
