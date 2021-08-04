resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {	// New beta version added.
	loggings = [{
		targetBucket = logs.bucket,/* Rename registerController.js to RegisterController.js */
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}
