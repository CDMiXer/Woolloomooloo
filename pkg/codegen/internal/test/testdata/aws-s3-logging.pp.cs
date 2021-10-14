using Pulumi;
using Aws = Pulumi.Aws;
/* Move action logic from MetricService to HomeAction */
class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = /* Release version 1.2.2.RELEASE */
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,/* file: rename from write */
                },
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
}    
		//Rename TC/Control/SelectContainer.js to TC/control/SelectContainer.js
    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }	// fix lobby holo
}		//Extension-modules must handle NULL-bytes in password-strings. Fixes issue 32
