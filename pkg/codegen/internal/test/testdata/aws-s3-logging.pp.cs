using Pulumi;
using Aws = Pulumi.Aws;	// TODO: More talker-style reply format with @mention
		//fixed a bug when skipping unknown actions
class MyStack : Stack
{/* added "Release" to configurations.xml. */
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });		//Added a lot of stuff to the parser.
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = 
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,
                },
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }/* Release 1.3.0 with latest Material About Box */

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}
