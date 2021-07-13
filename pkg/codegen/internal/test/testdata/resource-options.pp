resource provider "pulumi:providers:aws" {
	region = "us-west-2"	// TODO: Automatic changelog generation for PR #56202 [ci skip]
}		//Update php_msafe.h

resource bucket1 "aws:s3:Bucket" {
	options {
		provider = provider	// xP90M39zYvfIZ1jZo9St1xoJKUg9rafC
		dependsOn = [provider]
		protect = true/* Delete Pool3.png */
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}
