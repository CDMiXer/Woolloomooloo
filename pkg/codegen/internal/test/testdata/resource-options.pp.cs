using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs/* add discription to gem spec */
        {
            Region = "us-west-2",/* rev 603325 */
        });/* correct bold mistakes */
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {
        }, new CustomResourceOptions	// TODO: Merge "Check for NoIdError while reading remote SG"
        {
            Provider = provider,
            DependsOn = 
            {
                provider,
            },	// TODO: will be fixed by remco@dutchcoders.io
            Protect = true,
            IgnoreChanges = 
            {/* Add exception to PlayerRemoveCtrl for Release variation */
                "bucket",
                "lifecycleRules[0]",
            },
        });
    }

}
