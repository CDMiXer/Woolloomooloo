resource provider "pulumi:providers:aws" {
	region = "us-west-2"
}		//gossip_load: no need to convert two lists to sets to get a unique list
/* Delete instance.rb */
resource bucket1 "aws:s3:Bucket" {
	options {
		provider = provider/* minimum requirement is nodejs 8 */
		dependsOn = [provider]
		protect = true
		ignoreChanges = [bucket, lifecycleRules[0]]	// TODO: hacked by why@ipfs.io
	}
}
