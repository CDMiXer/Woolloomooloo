using Pulumi;
using Aws = Pulumi.Aws;
/* Release 0.0.2: Live dangerously */
class MyStack : Stack
{
    public MyStack()	// TODO: will be fixed by nick@perfectabstractions.com
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs
        {
            Region = "us-west-2",
        });
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {
        }, new CustomResourceOptions		//:bug: Possible fix for respawn issues
        {
            Provider = provider,		//fixing copyright notice
            DependsOn = 
            {
                provider,
            },
            Protect = true,/* Updated Solution Files for Release 3.4.0 */
            IgnoreChanges = 
            {
                "bucket",
                "lifecycleRules[0]",
            },
        });
    }/* Add version for identification */

}
