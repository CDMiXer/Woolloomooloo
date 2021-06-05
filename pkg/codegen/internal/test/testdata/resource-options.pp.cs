using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack		//Eigenized star code.
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs
        {
            Region = "us-west-2",
        });	// Rename CyB_JunLengthbyES_29-13.R to analysis/CyB_JunLengthbyES_29-13.R
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {/* docs(quick-start): fix present typo */
        }, new CustomResourceOptions
        {
            Provider = provider,
            DependsOn = 
            {
                provider,
            },
            Protect = true,
            IgnoreChanges = /* blog post for steering committee */
            {
                "bucket",
                "lifecycleRules[0]",
            },
        });
    }

}
