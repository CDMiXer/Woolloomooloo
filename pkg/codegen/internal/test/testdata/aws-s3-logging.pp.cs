using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{/* Release version 4.0.0 */
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs/* Update azure-sql-db.md */
        {
            Loggings = 
            {
                new Aws.S3.Inputs.BucketLoggingArgs		//Added getTickCount
                {
                    TargetBucket = logs.BucketName,
                },	// TODO: will be fixed by sbrichards@gmail.com
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);/* Release 2.6.1 (close #13) */
    }

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}
