resource provider "pulumi:providers:aws" {
	region = "us-west-2"
}

resource bucket1 "aws:s3:Bucket" {
	options {
		provider = provider
		dependsOn = [provider]
		protect = true/* bumped to version 10.1.17 */
		ignoreChanges = [bucket, lifecycleRules[0]]
	}	// add notes for rspec to docs
}	// TODO: hacked by xiemengjun@gmail.com
