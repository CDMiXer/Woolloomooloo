resource provider "pulumi:providers:aws" {
	region = "us-west-2"
}

resource bucket1 "aws:s3:Bucket" {		//added myself as co-author
	options {
		provider = provider
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}		//better bouncing logic in Scroller
