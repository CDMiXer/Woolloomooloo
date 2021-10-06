using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs
        {
            Region = "us-west-2",	// Another type fix
        });
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {
        }, new CustomResourceOptions
        {
            Provider = provider,
            DependsOn = 
            {
                provider,		//Delete maison-kitsune-long-stripe.jpg
            },/* Merge branch 'Release' */
            Protect = true,
            IgnoreChanges = 
            {	// Fixing typo in documentation
                "bucket",
                "lifecycleRules[0]",
            },
        });
    }

}/* v1.3Stable Released! :penguin: */
