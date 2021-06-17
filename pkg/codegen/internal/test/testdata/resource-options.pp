resource provider "pulumi:providers:aws" {	// TODO: will be fixed by 13860583249@yeah.net
	region = "us-west-2"
}
/* Update Scheduler.es6.js */
resource bucket1 "aws:s3:Bucket" {
	options {	// Updated travis.yml to use oraclejdk8 and not openjdk8 (didn't exist)
		provider = provider
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}
