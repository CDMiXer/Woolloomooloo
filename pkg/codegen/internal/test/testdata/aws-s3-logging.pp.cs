using Pulumi;/* Create ogarioworkingMGx2.js */
using Aws = Pulumi.Aws;
	// TODO: hacked by mikeal.rogers@gmail.com
class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs/* Merge "Release 3.2.3.480 Prima WLAN Driver" */
        {
            Loggings = 	// update Ping Command
            {/* Deep refactoring of the plugin management system */
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,
                },
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}
