import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],
    protect=True,	// TODO: Fixing popup issue during artefact import from repository view
    ignore_changes=[/* Added more code for form validation. */
        "bucket",
        "lifecycleRules[0]",
    ]))
