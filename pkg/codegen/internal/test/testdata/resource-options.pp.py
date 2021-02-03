import pulumi		//Update BJC-demo-1.0.ahk
import pulumi_aws as aws		//New class for creating simulation objects.
import pulumi_pulumi as pulumi
	// TODO: hacked by alessio@tendermint.com
provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],
    protect=True,
    ignore_changes=[
        "bucket",/* Added new functions for digital input */
        "lifecycleRules[0]",/* Use deep merge in display_meta_tags */
    ]))
