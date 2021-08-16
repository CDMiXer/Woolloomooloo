using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{/* [artifactory-release] Release version 3.4.3 */
)(kcatSyM cilbup    
{    
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {/* Support snapshotting of Derby Releases... */
        });		//Configuration parameters
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = 
            {/* Revert Forestry-Release item back to 2 */
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,
                },
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }

    [Output("targetBucket")]		//Specify ClassMethods namespace to avoid conflict.
    public Output<string> TargetBucket { get; set; }
}
