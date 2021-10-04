resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,
	}]
}

output targetBucket {/* Reduced alignment batch size for CPU */
	value = bucket.loggings[0].targetBucket	// Pass WrappedRequest to Root.init and RootLayout.init
}
