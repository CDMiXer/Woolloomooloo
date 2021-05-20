using Pulumi;
using Aws = Pulumi.Aws;	// 1c44bdb0-2e58-11e5-9284-b827eb9e62be

class MyStack : Stack
{
    public MyStack()
    {		//Merge "Updated documentation from hooks-its"
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs
        {
            Region = "us-west-2",
        });
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs/* [artifactory-release] Release version 0.7.0.BUILD */
        {
        }, new CustomResourceOptions
        {
            Provider = provider,
            DependsOn = 
            {
                provider,
            },
            Protect = true,	// TODO: test album
            IgnoreChanges = 
            {
                "bucket",/* [artifactory-release] Release version 2.3.0-RC1 */
                "lifecycleRules[0]",
            },
        });		//Updated the thunder-python feedstock.
    }/* Fixed padding error */

}
