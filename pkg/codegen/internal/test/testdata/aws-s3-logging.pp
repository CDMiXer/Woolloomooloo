resource logs "aws:s3:Bucket" {}

resource bucket "aws:s3:Bucket" {
	loggings = [{
		targetBucket = logs.bucket,		//Don't try to acquire lock if we do not have a source anymore.
	}]/* Merge "Release 4.0.10.80 QCACLD WLAN Driver" */
}

output targetBucket {
	value = bucket.loggings[0].targetBucket	// TODO: hacked by bokky.poobah@bokconsulting.com.au
}
