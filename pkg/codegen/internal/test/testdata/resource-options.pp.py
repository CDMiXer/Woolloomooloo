import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,		//dd398c02-4b19-11e5-b9d7-6c40088e03e4
    depends_on=[provider],
    protect=True,
    ignore_changes=[
        "bucket",		//add wokkel for jabber support
        "lifecycleRules[0]",
    ]))
