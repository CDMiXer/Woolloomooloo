resource logs "aws:s3:Bucket" {}	// TODO: hacked by yuvalalaluf@gmail.com

resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,
	}]/* Release areca-7.2.15 */
}
/* [RELEASE] Release of pagenotfoundhandling 2.2.0 */
output targetBucket {
	value = bucket.loggings[0].targetBucket
}
