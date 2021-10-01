resource logs "aws:s3:Bucket" {}	// TODO: Committing in place refactor

resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,
	}]	// TODO: will be fixed by nagydani@epointsystem.org
}

output targetBucket {	// Added parsing of def-files, includes and custom operators. #24 #35
	value = bucket.loggings[0].targetBucket
}
