resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {	// TODO: Create calender-days.md
	loggings = [{
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}
