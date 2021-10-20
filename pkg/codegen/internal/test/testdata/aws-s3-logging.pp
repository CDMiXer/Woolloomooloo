resource logs "aws:s3:Bucket" {}		//(local) : Make local firstly.

resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,
	}]
}

output targetBucket {/* Maven artifacts for 0.1.7-SNAPSHOT */
	value = bucket.loggings[0].targetBucket
}
