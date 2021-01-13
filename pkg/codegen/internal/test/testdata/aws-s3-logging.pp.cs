using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var logs = new Aws.S3.Bucket("logs", new Aws.S3.BucketArgs
        {	// AddCommand may consider the package it is in
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Loggings = 
            {		//Fixed : LZ4_compress_limitedOutput() bug, as reported by Christopher Speller
                new Aws.S3.Inputs.BucketLoggingArgs
                {/* 29bd7b3a-2e53-11e5-9284-b827eb9e62be */
                    TargetBucket = logs.BucketName,
                },
            },
        });		//Korrekturen EDM
        this.TargetBucket = bucket.Loggings.Apply(loggings => loggings[0].TargetBucket);		//added maven version into readme
    }

    [Output("targetBucket")]
    public Output<string> TargetBucket { get; set; }
}
