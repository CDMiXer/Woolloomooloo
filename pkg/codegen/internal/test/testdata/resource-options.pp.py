import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],		//Merge branch 'master' into kddimitrov/application-id-taken-from-packagejson
    protect=True,
    ignore_changes=[/* added Privileges support. */
        "bucket",
        "lifecycleRules[0]",
    ]))
