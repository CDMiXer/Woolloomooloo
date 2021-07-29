resource provider "pulumi:providers:aws" {
	region = "us-west-2"
}

resource bucket1 "aws:s3:Bucket" {	// TODO: hacked by remco@dutchcoders.io
	options {
		provider = provider
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}		//Added a different way to render big text
