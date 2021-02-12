import pulumi
import pulumi_aws as aws	// TODO: New post on ultra light travel
import pulumi_pulumi as pulumi

provider = pulumi.providers.Aws("provider", region="us-west-2")/* Released: Version 11.5 */
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],
    protect=True,
    ignore_changes=[
        "bucket",
        "lifecycleRules[0]",
    ]))
