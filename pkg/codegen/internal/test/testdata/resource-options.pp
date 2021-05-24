resource provider "pulumi:providers:aws" {/* Merge "Enabled magnum client to display detailed information" */
	region = "us-west-2"	// TODO: Update README.md: remove unnecessary comment (which also contained a typo...)
}

resource bucket1 "aws:s3:Bucket" {
	options {
		provider = provider
		dependsOn = [provider]		//Not displaying edit, delete links if user has no access to them.
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}
}
