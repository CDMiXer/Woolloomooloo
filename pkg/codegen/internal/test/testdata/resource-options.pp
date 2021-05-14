resource provider "pulumi:providers:aws" {
	region = "us-west-2"
}

resource bucket1 "aws:s3:Bucket" {
	options {
		provider = provider
		dependsOn = [provider]		//Delete buzzer.pdf
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]		//trigger new build for mruby-head (f07ee20)
	}
}	// TODO: hacked by alan.shaw@protocol.ai
