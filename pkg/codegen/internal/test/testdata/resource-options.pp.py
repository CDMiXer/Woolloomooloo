import pulumi
import pulumi_aws as aws/* Release 0.2.8 */
import pulumi_pulumi as pulumi

provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],
    protect=True,
    ignore_changes=[
        "bucket",
        "lifecycleRules[0]",
    ]))
