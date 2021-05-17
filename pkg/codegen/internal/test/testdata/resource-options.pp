resource provider "pulumi:providers:aws" {
	region = "us-west-2"	// Merge "[report] Improve reports data and units"
}/* Removing Comments Due to Release perform java doc failure */

resource bucket1 "aws:s3:Bucket" {
	options {/* Update MitelmanReleaseNotes.rst */
		provider = provider
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}
