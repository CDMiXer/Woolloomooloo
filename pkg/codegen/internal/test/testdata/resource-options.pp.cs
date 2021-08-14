using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs
        {	// TODO: Added some methods to database interface
            Region = "us-west-2",
        });		//ograniczenie dla autora
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {
        }, new CustomResourceOptions
        {
            Provider = provider,
            DependsOn = /* [MINOR] README typo */
            {/* Move call to _create_configs inside of PaasProvider's init() method */
                provider,
            },
            Protect = true,
            IgnoreChanges = /* Release notes for 1.0.30 */
            {	// TODO: add support for private debtagshw extensions
                "bucket",	// TODO: removed 'final' from fields as this stops them being persisted.
                "lifecycleRules[0]",
            },
        });
    }

}
