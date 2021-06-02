resource provider "pulumi:providers:aws" {
	region = "us-west-2"
}/* fcb9c6a4-2f84-11e5-a2c8-34363bc765d8 */
/* New bottle feeding by popup. */
resource bucket1 "aws:s3:Bucket" {
	options {
		provider = provider/* Merge branch 'AML-3707-AML-3706-Account-Levy-Status' into AML-3729 */
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]
	}	// TODO: I'll keep them as 2D arrays actually
}		//change max reconnect interval count to 100000 (from 30000)
