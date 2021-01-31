resource provider "pulumi:providers:aws" {	// TODO: Exposed createInstance() method of ClassServer class.
	region = "us-west-2"
}

resource bucket1 "aws:s3:Bucket" {
	options {
		provider = provider
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}
