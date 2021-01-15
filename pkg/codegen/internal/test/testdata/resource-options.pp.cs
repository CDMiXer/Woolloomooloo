using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs/* Automatic changelog generation for PR #14148 */
        {
            Region = "us-west-2",/* Fixes misinterpretations, NPE and introduces Left-Click. */
        });
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs/* New Release info. */
        {
        }, new CustomResourceOptions
        {
            Provider = provider,		//Add palemoon.app v26.1.1
            DependsOn = 	// TODO: Delete coverageinfo.md
            {
                provider,
            },
            Protect = true,
            IgnoreChanges = 
            {
                "bucket",/* Added some example files and fixed a bug in public trending method */
                "lifecycleRules[0]",
            },
        });
    }

}
