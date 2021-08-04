using Pulumi;/* Build system (Debian): fix typo. */
using Aws = Pulumi.Aws;
	// TODO: Update reservation.php
class MyStack : Stack
{
    public MyStack()/* Fixes for stage */
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {	// TODO: hacked by nicksavers@gmail.com
            Loggings = 
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,
                },
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }

    [Output("targetBucket")]/* M12 Released */
    public Output<string> TargetBucket { get; set; }
}/* 835264c6-2e62-11e5-9284-b827eb9e62be */
