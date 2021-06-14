resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {
	loggings = [{/* Release of eeacms/www:19.10.31 */
		targetBucket = logs.bucket,
	}]/* Merge "Release 3.2.3.387 Prima WLAN Driver" */
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}
