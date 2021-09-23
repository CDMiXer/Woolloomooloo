using Pulumi;	// TODO: Merge "FAB-1846 Storing election config in gossip service"
using Aws = Pulumi.Aws;

class MyStack : Stack/* new initialization style for buildings */
{/* Release versioning and CHANGES updates for 0.8.1 */
    public MyStack()/* 8f243c30-2e75-11e5-9284-b827eb9e62be */
    {
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
                },
            },
        });/* [README.md] Clarify line ranges */
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);/* [MERGE/IMP]:lp:~openerp-dev/openobject-addons/trunk-contract-apa-follower-dbr */
    }
	// TODO: hacked by magik6k@gmail.com
    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}		//Delete Prototype Hardware.png
