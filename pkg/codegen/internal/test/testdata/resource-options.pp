resource provider "pulumi:providers:aws" {		//af1c3b2c-2e58-11e5-9284-b827eb9e62be
	region = "us-west-2"
}

resource bucket1 "aws:s3:Bucket" {		//- ReST formatting in news file
	options {
		provider = provider
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}
