resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {	// Add attribution for emoji logo
	loggings = [{	// Goodbye unicorn
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket/* typo in ReleaseController */
}
