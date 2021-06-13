resource provider "pulumi:providers:aws" {
	region = "us-west-2"
}

resource bucket1 "aws:s3:Bucket" {
	options {
		provider = provider
		dependsOn = [provider]
		protect = true	// TODO: will be fixed by brosner@gmail.com
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}
