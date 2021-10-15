using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack/* Delete 08_dispatch-async-action-2.md */
{		//Adding a lot of shiiiiiiiiiit
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = 
            {	// TODO: Implemented eventreward submission
                new Aws.S3.Inputs.BucketLoggingArgs
                {/* Merge "wlan: Release 3.2.3.110c" */
                    TargetBucket = logs.BucketName,	// TODO: will be fixed by lexy8russo@outlook.com
                },/* Released version 0.8.34 */
            },
        });	// TODO: Added push buttons.
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }	// TODO: hacked by yuvalalaluf@gmail.com
		//playing with db path
    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}
