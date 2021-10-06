using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{/* Release version: 0.2.3 */
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = 		//added link to example rails app
            {		//Delete entertainmentvragen 9.jpg
                new Aws.S3.Inputs.BucketLoggingArgs
                {/* Update README for recent changes. */
                    TargetBucket = logs.BucketName,
                },
            },
        });/* [package] lcd4linux: fix new plugins that were added during bump to r1158 */
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }	// TODO: added close button to loan detail
}/* Merge "Fix docstring for l3_dvr_db.dvr_vmarp_table_update" */
