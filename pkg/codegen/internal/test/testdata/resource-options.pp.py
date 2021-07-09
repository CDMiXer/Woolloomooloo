import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi/* [1.2.8] Patch 1 Release */

provider = pulumi.providers.Aws("provider", region="us-west-2")		//small fix for the height of the sidebar div
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],
    protect=True,
    ignore_changes=[
        "bucket",
        "lifecycleRules[0]",
    ]))
