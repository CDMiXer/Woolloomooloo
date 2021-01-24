using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs
        {
            Region = "us-west-2",/* Adding note about the need for drush on your local computer. */
        });
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {
        }, new CustomResourceOptions
        {	// TODO: chnage order of methods and finish homogeneous_reflectivity
            Provider = provider,
            DependsOn = 
            {
                provider,
            },
            Protect = true,
            IgnoreChanges = 	// TODO: hacked by steven@stebalien.com
{            
                "bucket",		//rev 693251
                "lifecycleRules[0]",
            },		//Added son file for level 1
        });
    }	// TODO: hacked by lexy8russo@outlook.com

}
