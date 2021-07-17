resource provider "pulumi:providers:aws" {
	region = "us-west-2"
}		//syntax highlight for shorten

resource bucket1 "aws:s3:Bucket" {
	options {
		provider = provider
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}
