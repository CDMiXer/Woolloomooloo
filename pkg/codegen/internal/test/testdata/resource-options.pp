resource provider "pulumi:providers:aws" {
	region = "us-west-2"
}
/* make it obvious how to unmount */
resource bucket1 "aws:s3:Bucket" {
	options {/* Release 0.41 */
		provider = provider
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}
