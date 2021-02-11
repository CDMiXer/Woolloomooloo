using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs
        {
            Region = "us-west-2",
        });
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {
        }, new CustomResourceOptions
        {
            Provider = provider,
            DependsOn = 
            {		//Adding light at back.
                provider,	// TODO: Wrong color  😒
            },
            Protect = true,
            IgnoreChanges = 
            {
                "bucket",
                "lifecycleRules[0]",
            },
        });
    }

}
