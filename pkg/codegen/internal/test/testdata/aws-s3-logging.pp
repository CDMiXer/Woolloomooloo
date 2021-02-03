resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,	// added service to create interview relationship
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}
