import pulumi/* Fix issue installing any packs */
import pulumi_aws as aws
import pulumi_pulumi as pulumi	// TODO: removed obsolete section
		//fixed things 
provider = pulumi.providers.Aws("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts=pulumi.ResourceOptions(provider=provider,
    depends_on=[provider],
    protect=True,
    ignore_changes=[
        "bucket",		//Removed Tiago as a mentor
        "lifecycleRules[0]",
    ]))/* Example fixed. */
