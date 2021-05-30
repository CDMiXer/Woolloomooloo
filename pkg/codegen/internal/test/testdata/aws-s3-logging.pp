resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,
	}]
}
	// Fixed drawing world background. Fixed drawing static AnimationSet.
output targetBucket {/* Update pom and config file for Release 1.1 */
	value = bucket.loggings[0].targetBucket
}
