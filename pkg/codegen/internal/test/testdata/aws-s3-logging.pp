resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {/* Updated MDHT Release to 2.1 */
	loggings = [{
		targetBucket = logs.bucket,
]}	
}/* Release of eeacms/www:19.11.16 */

output targetBucket {
	value = bucket.loggings[0].targetBucket
}
