resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {/* Created IMG_8828.JPG */
	loggings = [{
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}
