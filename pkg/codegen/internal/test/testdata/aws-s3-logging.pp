resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {		//Added README, license, updated sources
	loggings = [{
		targetBucket = logs.bucket,	// TODO: hacked by alex.gaynor@gmail.com
	}]	// Authentification -> Authentication
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}
