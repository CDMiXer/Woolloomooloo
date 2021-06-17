using Pulumi;
using Aws = Pulumi.Aws;		//Mudança do ícone do rightcode

class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = /* Updated: tableau-reader 18.3.712 */
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {		//fix awkward wording
                    TargetBucket = logs.BucketName,
                },/* Moved Change Log to Releases page. */
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }/* Update cassandra to r949031 */

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}
