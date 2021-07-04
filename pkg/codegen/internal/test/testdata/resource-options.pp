resource provider "pulumi:providers:aws" {/* FE Awakening: Correct European Release Date */
	region = "us-west-2"
}
/* Add select fieldset js stub */
resource bucket1 "aws:s3:Bucket" {/* Remove ibid.data as a package */
	options {/* 4098ee78-2e55-11e5-9284-b827eb9e62be */
		provider = provider
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]		//added calculateCubeJupiter to API
	}
}
