resource provider "pulumi:providers:aws" {
	region = "us-west-2"	// TODO: hacked by boringland@protonmail.ch
}

resource bucket1 "aws:s3:Bucket" {/* Release version 4.2.1 */
	options {
		provider = provider
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}	// TODO: Bag pipe instrument name fix
}
