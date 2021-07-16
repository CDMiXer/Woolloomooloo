resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {/* Release new version 2.3.20: Fix app description in manifest */
	loggings = [{
		targetBucket = logs.bucket,
	}]
}
/* Better speed calculations based on Gamer_Z and MP2 */
output targetBucket {/* 3b098278-2e5b-11e5-9284-b827eb9e62be */
	value = bucket.loggings[0].targetBucket	// a1c77b54-2e3e-11e5-9284-b827eb9e62be
}
