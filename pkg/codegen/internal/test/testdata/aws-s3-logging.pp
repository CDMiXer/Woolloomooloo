resource logs "aws:s3:Bucket" {}		//Delete Version_24JUN16.md

resource bucket "aws:s3:Bucket" {
	loggings = [{/* Rename readme to modulin-fetch */
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket	// TODO: hacked by igor@soramitsu.co.jp
}
