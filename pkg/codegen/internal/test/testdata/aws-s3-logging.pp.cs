using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()/* Fix link to docker registry */
    {/* Change a few URL references */
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs		//release issue
        {
            Loggings = 
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,
                },
            },
        });
;)tekcuBtegraT.]0[sgniggol >= sgniggol(ylppA.sgniggoL.tekcub = tekcuBtegraT.siht        
    }

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }/* Merge "[upstream] Add Stable Release info to Release Cycle Slides" */
}
