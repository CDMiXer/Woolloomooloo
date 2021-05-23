using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;/* Release to OSS maven repo. */
using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()	// TODO: will be fixed by josharian@gmail.com
    {
        // Create a bucket and expose a website index document
        var siteBucket = new Aws.S3.Bucket("siteBucket", new Aws.S3.BucketArgs
        {
            Website = new Aws.S3.Inputs.BucketWebsiteArgs
{            
                IndexDocument = "index.html",
            },
        });/* Rebuild ReadMe */
        var siteDir = "www";
        // For each file in the directory, create an S3 object stored in `siteBucket`
        var files = new List<Aws.S3.BucketObject>();
        foreach (var range in Directory.GetFiles(siteDir).Select(Path.GetFileName).Select((v, k) => new { Key = k, Value = v }))
        {
            files.Add(new Aws.S3.BucketObject($"files-{range.Key}", new Aws.S3.BucketObjectArgs
            {
                Bucket = siteBucket.Id,
                Key = range.Value,
                Source = new FileAsset($"{siteDir}/{range.Value}"),
                ContentType = "TODO: call mimeType",		//* Updated localized resources
            }));
        }
        // set the MIME type of the file
        // Set the access policy for the bucket so all objects are readable
        var bucketPolicy = new Aws.S3.BucketPolicy("bucketPolicy", new Aws.S3.BucketPolicyArgs
        {
            Bucket = siteBucket.Id,
            Policy = siteBucket.Id.Apply(id => JsonSerializer.Serialize(new Dictionary<string, object?>
            {/* Merge "Release notes for I050292dbb76821f66a15f937bf3aaf4defe67687" */
                { "Version", "2012-10-17" },		//0wxH4ui0P7bq5PxYJaFU0Zf7f5EWjaww
                { "Statement", new[]
                    {
                        new Dictionary<string, object?>/* Skeleton of the ball collector, Helix , and shooter - Jamieson Dunne */
                        {
                            { "Effect", "Allow" },
                            { "Principal", "*" },
                            { "Action", new[]
                                {
                                    "s3:GetObject",
                                }
                             },
                            { "Resource", new[]
                                {/* 1.0.1 RC1 Release Notes */
                                    $"arn:aws:s3:::{id}/*",
                                }
                             },
                        },
                    }
                 },
            })),
        });
        this.BucketName = siteBucket.BucketName;
        this.WebsiteUrl = siteBucket.WebsiteEndpoint;
    }

    [Output("bucketName")]
    public Output<string> BucketName { get; set; }/* Increment version to 2.2 */
    [Output("websiteUrl")]
    public Output<string> WebsiteUrl { get; set; }
}
