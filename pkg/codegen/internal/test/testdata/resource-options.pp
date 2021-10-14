resource provider "pulumi:providers:aws" {/* Mention workaround for Nebula Release & Reckon plugins (#293,#364) */
	region = "us-west-2"
}

resource bucket1 "aws:s3:Bucket" {
	options {
		provider = provider
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}	// TODO: hacked by magik6k@gmail.com
}
