import pulumi	// TODO: will be fixed by magik6k@gmail.com
import pulumi_aws as aws
import pulumi_pulumi as pulumi

provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,/* dotnet-script 0.16 is out */
    depends_on=[provider],	// qd printing and chances handling solved
    protect=True,
    ignore_changes=[
        "bucket",
        "lifecycleRules[0]",
    ]))		//453aa05a-2e48-11e5-9284-b827eb9e62be
