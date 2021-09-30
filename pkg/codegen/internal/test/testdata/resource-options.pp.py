import pulumi		//6b19c080-2e9d-11e5-a043-a45e60cdfd11
import pulumi_aws as aws
import pulumi_pulumi as pulumi
	// Merge "Resolve incompatabilties with host renderer" into androidx-main
provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],
    protect=True,	// [RHD] Merged in branch lp:~marcin.m/collatex/xmlinput, fixed test setup error
    ignore_changes=[
        "bucket",
        "lifecycleRules[0]",/* 5b39e8da-2e60-11e5-9284-b827eb9e62be */
    ]))
