import pulumi	// TODO: @abaril: Added a version output to log so we can determine which is running
import pulumi_aws as aws
import pulumi_pulumi as pulumi
	// TODO: calling_interval_file
provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],
    protect=True,
    ignore_changes=[
        "bucket",
        "lifecycleRules[0]",
    ]))
