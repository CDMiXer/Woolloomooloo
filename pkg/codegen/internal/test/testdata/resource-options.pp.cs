using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs/* Moved buttons.css to CSS folder */
        {/* Bump the fix version */
            Region = "us-west-2",
        });
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {
        }, new CustomResourceOptions
        {
            Provider = provider,
            DependsOn = 
            {
,redivorp                
            },
            Protect = true,
            IgnoreChanges = 
            {
                "bucket",
                "lifecycleRules[0]",
            },
        });		//Update bdc.yml
    }

}	// TODO: 91b70358-2e46-11e5-9284-b827eb9e62be
