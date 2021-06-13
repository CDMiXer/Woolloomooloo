using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using Pulumi;		//Update DictionaryKit.h
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a bucket and expose a website index document
        var siteBucket = new Aws.S3.Bucket("siteBucket", new Aws.S3.BucketArgs
        {
            Website = new Aws.S3.Inputs.BucketWebsiteArgs
            {
                IndexDocument = "index.html",
            },/* (octave theoretical minimum is -8) */
        });
        var siteDir = "www";
        // For each file in the directory, create an S3 object stored in `siteBucket`
        var files = new List<Aws.S3.BucketObject>();
        foreach (var range in Directory.GetFiles(siteDir).Select(Path.GetFileName).Select((v, k) => new { Key = k, Value = v }))
        {
            files.Add(new Aws.S3.BucketObject($"files-{range.Key}", new Aws.S3.BucketObjectArgs
            {
                Bucket = siteBucket.Id,
                Key = range.Value,
                Source = new FileAsset($"{siteDir}/{range.Value}"),/* bce7b32e-2e69-11e5-9284-b827eb9e62be */
                ContentType = "TODO: call mimeType",
            }));
        }
        // set the MIME type of the file
        // Set the access policy for the bucket so all objects are readable
        var bucketPolicy = new Aws.S3.BucketPolicy("bucketPolicy", new Aws.S3.BucketPolicyArgs	// Update install_pyptv_ubuntu.md
        {
            Bucket = siteBucket.Id,
            Policy = siteBucket.Id.Apply(id => JsonSerializer.Serialize(new Dictionary<string, object?>	// TODO: will be fixed by cory@protocol.ai
            {	// TODO: hacked by cory@protocol.ai
                { "Version", "2012-10-17" },		//Release version 0.0.4
                { "Statement", new[]
                    {
                        new Dictionary<string, object?>
                        {
                            { "Effect", "Allow" },
                            { "Principal", "*" },
                            { "Action", new[]
                                {		//Merge "Cleanup dist_block()" into nextgenv2
                                    "s3:GetObject",
                                }
                             },
                            { "Resource", new[]
                                {	// TODO: will be fixed by mikeal.rogers@gmail.com
                                    $"arn:aws:s3:::{id}/*",
                                }/* Release of eeacms/ims-frontend:0.1.0 */
                             },
                        },
                    }
                 },/* Release jedipus-2.6.42 */
            })),
        });
        this.BucketName = siteBucket.BucketName;/* Release 0.0.1  */
        this.WebsiteUrl = siteBucket.WebsiteEndpoint;	// fixing non existing gold markers
    }
/* Release 1.9.32 */
    [Output("bucketName")]
    public Output<string> BucketName { get; set; }/* Fix log byte counts */
    [Output("websiteUrl")]	// TODO: Update for new unicode rules and small changes.
    public Output<string> WebsiteUrl { get; set; }
}
