using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {	// Added the beginings of a scene graph.
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {/* Added array_filter. */
            Loggings = 
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,		//reorganization, parser work
                },
            },/* [ADD] Beta and Stable Releases */
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}	// TODO: 494cfa21-2d48-11e5-befb-7831c1c36510
