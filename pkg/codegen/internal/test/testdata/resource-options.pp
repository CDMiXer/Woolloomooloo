resource provider "pulumi:providers:aws" {
	region = "us-west-2"
}

resource bucket1 "aws:s3:Bucket" {/* Release v0.92 */
	options {
		provider = provider
		dependsOn = [provider]		//Edit Progress Report + BAB 3.2
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}/* add freeze dry to lapras */
}
