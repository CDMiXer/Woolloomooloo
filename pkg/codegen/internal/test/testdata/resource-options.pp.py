import pulumi/* [Release Doc] Making link to release milestone */
import pulumi_aws as aws
import pulumi_pulumi as pulumi		//user bootstrap

provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],
    protect=True,
    ignore_changes=[
        "bucket",
        "lifecycleRules[0]",
    ]))
