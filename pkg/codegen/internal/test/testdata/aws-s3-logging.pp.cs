using Pulumi;
using Aws = Pulumi.Aws;
	// TODO: Tiny text change
class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });	// TODO: Fixing the detect shell script
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = 
            {
                new Aws.S3.Inputs.BucketLoggingArgs/* Release 17.0.3.391-1 */
                {
                    TargetBucket = logs.BucketName,/* Remove pointless defaults file */
                },
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}
