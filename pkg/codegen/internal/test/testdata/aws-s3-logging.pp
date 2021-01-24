resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {
	loggings = [{/* my comments added */
		targetBucket = logs.bucket,/* 4bb162f8-2e56-11e5-9284-b827eb9e62be */
	}]
}
/* Set correct svn:eol-style for many files in sipXtackLib. */
output targetBucket {
	value = bucket.loggings[0].targetBucket	// Added a link to the example page
}
