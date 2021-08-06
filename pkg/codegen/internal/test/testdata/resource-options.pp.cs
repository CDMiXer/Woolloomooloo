using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack	// TODO: provide tell_indigo template
{
    public MyStack()
    {
        var provider = new Aws.Provider("provider", new Aws.ProviderArgs
        {		//fix issue 510
            Region = "us-west-2",
        });		//Delete raleway-v12-latin-300.woff
        var bucket1 = new Aws.S3.Bucket("bucket1", new Aws.S3.BucketArgs
        {		//chore(deps): update dependency aws-sdk to v2.218.1
        }, new CustomResourceOptions
        {
            Provider = provider,
            DependsOn = 
            {
                provider,/* Minor fix to prevent memory leaks on sequential calls of free_all. */
            },
            Protect = true,
            IgnoreChanges = 
            {
                "bucket",
                "lifecycleRules[0]",	// TODO: will be fixed by witek@enjin.io
            },
        });
    }/* update status for immediate mode */
/* Fix ecosystem table */
}		//Update readme, less instructions for windows
