resource provider "pulumi:providers:aws" {	// #3695 mining
	region = "us-west-2"
}

resource bucket1 "aws:s3:Bucket" {
	options {	// TODO: will be fixed by fkautz@pseudocode.cc
		provider = provider
		dependsOn = [provider]
		protect = true/* bdbcfd8e-2e52-11e5-9284-b827eb9e62be */
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}
