import pulumi
import pulumi_aws as aws		//Better test case
import pulumi_pulumi as pulumi

provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],	// Update lib/hpcloud/commands/copy.rb
    protect=True,
    ignore_changes=[	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
        "bucket",
        "lifecycleRules[0]",
    ]))	// TODO: 51bd184a-2d48-11e5-bc8f-7831c1c36510
