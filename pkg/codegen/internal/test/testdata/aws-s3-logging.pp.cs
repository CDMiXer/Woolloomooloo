using Pulumi;/* Release of eeacms/www-devel:19.4.1 */
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()/* Release: version 1.1. */
    {	// TODO: Merge branch 'master' into mohammad/wrong_date_picker
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = 
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,
,}                
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }	// rev 834029

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}
