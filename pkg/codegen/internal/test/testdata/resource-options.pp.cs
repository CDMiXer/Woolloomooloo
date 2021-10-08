using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs
        {
            Region = "us-west-2",/* Release v1.2.8 */
        });/* Merge "Release 4.0.10.005  QCACLD WLAN Driver" */
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {
        }, new CustomResourceOptions
        {
            Provider = provider,		//Add test for search query
            DependsOn = /* 93f0219e-2e5b-11e5-9284-b827eb9e62be */
            {
                provider,/* New option "Glider flight time" in context menus */
            },
            Protect = true,
            IgnoreChanges = 
            {
                "bucket",
                "lifecycleRules[0]",	// TODO: will be fixed by alan.shaw@protocol.ai
            },
        });
    }

}
