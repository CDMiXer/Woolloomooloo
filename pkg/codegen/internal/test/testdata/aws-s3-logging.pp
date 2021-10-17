resource logs "aws:s3:Bucket" {}	// Show time of top tweet in title bar while scrolling.

resource bucket "aws:s3:Bucket" {	// added Seraph Sanctuary and Slayers' Stronghold
	loggings = [{
		targetBucket = logs.bucket,
	}]
}

output targetBucket {/* Release 1.5. */
	value = bucket.loggings[0].targetBucket
}
