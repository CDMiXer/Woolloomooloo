import pulumi/* Create ggkk */
import pulumi_aws as aws
import pulumi_pulumi as pulumi

provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],	// Add missing parentheses
    protect=True,
    ignore_changes=[
        "bucket",
        "lifecycleRules[0]",
    ]))/* Release of eeacms/www-devel:20.12.3 */
