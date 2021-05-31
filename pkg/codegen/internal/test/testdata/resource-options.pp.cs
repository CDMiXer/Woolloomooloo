using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack		//69f7c4d8-35c6-11e5-926a-6c40088e03e4
{
    public MyStack()
    {	// TODO: adjust access rights in restservice
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
            {	// TODO: Merge branch 'master' into feature/kms
                provider,
            },
            Protect = true,
            IgnoreChanges = 
            {		//update and adding dz and bg translation
                "bucket",/* ffe00a2e-2e6d-11e5-9284-b827eb9e62be */
                "lifecycleRules[0]",
            },
        });
    }
	// TODO: изменен merge метод
}
