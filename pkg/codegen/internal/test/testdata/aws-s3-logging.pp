resource logs "aws:s3:Bucket" {}/* Release notes updated to include checkbox + disable node changes */
		//Change phrasing on readme
resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,
	}]
}

output targetBucket {/* Release 3.0.2 */
	value = bucket.loggings[0].targetBucket	// TODO: hacked by cory@protocol.ai
}
