using Pulumi;/* Allow used profiles and queries to be deleted properly. */
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs/* Ignore files generated with the execution of the Maven Release plugin */
        {
            Loggings = 		//update Aardvark.Base.nuspec to v1.0.4
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,	// 4d7aaee6-2e4f-11e5-896c-28cfe91dbc4b
                },
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}
