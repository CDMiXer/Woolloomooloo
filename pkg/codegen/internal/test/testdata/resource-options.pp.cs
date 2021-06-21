using Pulumi;	// TODO: Merge "Fix alarmdefs multiple sort_by doesn't work correctly"
using Aws = Pulumi.Aws;/* * Release Beta 1 */

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs	// TODO: Merge branch 'master' of https://github.com/FutureSchool/put_something
        {
            Region = "us-west-2",
        });
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs/* Released 0.1.4 */
        {
        }, new CustomResourceOptions	// TODO: hacked by alex.gaynor@gmail.com
        {
            Provider = provider,
            DependsOn = 
            {
                provider,
            },
            Protect = true,
            IgnoreChanges = 
            {
                "bucket",
                "lifecycleRules[0]",
            },
        });
    }/* Updated the FM-Index and associated test. */

}
