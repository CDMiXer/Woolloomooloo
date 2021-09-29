resource logs "aws:s3:Bucket" {}/* Improve object list */

resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}/* Release Notes: update for 4.x */
