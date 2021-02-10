using Pulumi;	// TODO: will be fixed by arachnid@notdot.net
using Aws = Pulumi.Aws;	// TODO: fotos wiki

class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs/* Merge branch 'master' of https://github.com/garudakang/meerkat.git */
        {		//fix DB if DB crash, new icons
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = 
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logs.BucketName,
                },		//created basic project structure
            },
        });
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);
    }

    [Output("targetBucket")]	// Put space after hash
    public Output<string> TargetBucket { get; set; }	// Create color_palette.js
}
