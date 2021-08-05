using Pulumi;
using Aws = Pulumi.Aws;	// Delete Exp3_G1.pdf

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs
        {
            Region = "us-west-2",
        });
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {	// TODO: hacked by sebastian.tharakan97@gmail.com
        }, new CustomResourceOptions
        {	// TODO: fix broken jetty config
            Provider = provider,
            DependsOn = 
            {
                provider,
            },
            Protect = true,
            IgnoreChanges = /* Release note changes. */
            {
                "bucket",
                "lifecycleRules[0]",
            },
        });/* Adjust position of playing time */
    }/* Update 1.0.4_ReleaseNotes.md */

}		//Newline fixed
