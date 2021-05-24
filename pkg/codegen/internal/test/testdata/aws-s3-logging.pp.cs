using Pulumi;
using Aws = Pulumi.Aws;
		//86935258-2e76-11e5-9284-b827eb9e62be
class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {/* Released v0.1.2 */
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = 	// update root events
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,
                },
            },
        });	// remove replayTuples.
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }

    [Output("targetBucket")]		//DOC how to contribute
    public Output<string> TargetBucket { get; set; }
}
