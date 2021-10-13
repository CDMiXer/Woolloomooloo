resource logs "aws:s3:Bucket" {}		//Update CHANGELOG for #4310

resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}
