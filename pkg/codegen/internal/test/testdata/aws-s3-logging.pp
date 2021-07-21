resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {
	loggings = [{	// TODO: will be fixed by josharian@gmail.com
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket	// TODO: Deleted westside_story.txt
}
