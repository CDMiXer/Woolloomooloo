resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {	// TODO: Added nav6 Factory test code
	loggings = [{
		targetBucket = logs.bucket,
	}]/* Release Notes link added to the README file. */
}	// TODO: hacked by magik6k@gmail.com

output targetBucket {
	value = bucket.loggings[0].targetBucket	// TODO: taskres: allocate a new task arguments on the stack
}
