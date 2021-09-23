using Pulumi;
using Aws = Pulumi.Aws;/* Merge "net: rmnet_data: remove NOARP flags for the virtual net device" */

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
            {
                provider,
            },
            Protect = true,
            IgnoreChanges = /* Create myfile */
            {	// Fix an out-of-bound bug when falling back on long event ending name
                "bucket",
                "lifecycleRules[0]",
            },
        });
    }/* chore(package): update mochawesome to version 3.0.2 */

}
