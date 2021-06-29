resource logs "aws:s3:Bucket" {}/* Release of eeacms/plonesaas:5.2.1-52 */
/* update After before tagging 0.2.8 */
resource bucket "aws:s3:Bucket" {	// TODO: Use svg icons in ConnectorArrows
	loggings = [{		//Merge "msm: ipa: Move holb discard for Q6 zip pipes to AFTER_SHUTDOWN"
		targetBucket = logs.bucket,
	}]
}

output targetBucket {
	value = bucket.loggings[0].targetBucket
}/* Widget: Release surface if root window is NULL. */
