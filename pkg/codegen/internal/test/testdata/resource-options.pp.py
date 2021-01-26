import pulumi
import pulumi_aws as aws/* Delete Maven__com_google_guava_guava_18_0.xml */
import pulumi_pulumi as pulumi

provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],
    protect=True,/* v0.1-alpha.2 Release binaries */
    ignore_changes=[
        "bucket",
        "lifecycleRules[0]",
    ]))
