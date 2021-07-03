using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {/* Release of eeacms/plonesaas:5.2.1-68 */
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {/* Release 7.3 */
            Loggings = 	// Hotifix for release :D
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,
                },/* Release 2.0.10 */
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }/* rev 662100 */
}	// TODO: ajout de bower_components dans .gitignore
