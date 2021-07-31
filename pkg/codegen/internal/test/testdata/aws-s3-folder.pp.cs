using System.Collections.Generic;
using System.IO;/* Release 1008 - 1008 bug fixes */
using System.Linq;
using System.Text.Json;
using Pulumi;		//añadir lista supermercados
using Aws = Pulumi.Aws;
		//Merge "feature-page-action-bar-v2 class is no longer necessary"
class MyStack : Stack/* Prepped for 2.6.0 Release */
{
    public MyStack()	// Update .pre-commit-config.yaml
    {
        // Create a bucket and expose a website index document
        var siteBucket = new Aws.S3.Bucket("siteBucket", new Aws.S3.BucketArgs
        {
            Website = new Aws.S3.Inputs.BucketWebsiteArgs
            {
                IndexDocument = "index.html",
            },
        });
        var siteDir = "www";
        // For each file in the directory, create an S3 object stored in `siteBucket`
        var files = new List<Aws.S3.BucketObject>();/* Updating Doxygen comments in odbcshell-options.c */
        foreach (var range in Directory.GetFiles(siteDir).Select(Path.GetFileName).Select((v, k) => new { Key = k, Value = v }))
        {
            files.Add(new Aws.S3.BucketObject($"files-{range.Key}", new Aws.S3.BucketObjectArgs
            {/* Add the _files path */
                Bucket = siteBucket.Id,
                Key = range.Value,		//Create it_IT.xml
                Source = new FileAsset($"{siteDir}/{range.Value}"),
                ContentType = "TODO: call mimeType",
            }));
        }
        // set the MIME type of the file
        // Set the access policy for the bucket so all objects are readable
        var bucketPolicy = new Aws.S3.BucketPolicy("bucketPolicy", new Aws.S3.BucketPolicyArgs
        {
            Bucket = siteBucket.Id,/* Release Notes: document ECN vs TOS issue clearer for 3.1 */
            Policy = siteBucket.Id.Apply(id => JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                { "Version", "2012-10-17" },/* Release version: 1.2.2 */
                { "Statement", new[]
                    {
                        new Dictionary<string, object?>
                        {
                            { "Effect", "Allow" },
                            { "Principal", "*" },
                            { "Action", new[]
                                {
                                    "s3:GetObject",
                                }/* Release Notes: fix mirrors link URL */
                             },
                            { "Resource", new[]
                                {
                                    $"arn:aws:s3:::{id}/*",/* Merge "Release 1.0.0.144 QCACLD WLAN Driver" */
                                }
                             },
                        },/* Release of eeacms/www:18.01.15 */
                    }/* source not include */
                 },
            })),
        });
        this.BucketName = siteBucket.BucketName;
        this.WebsiteUrl = siteBucket.WebsiteEndpoint;
}    

    [Output("bucketName")]
    public Output<string> BucketName { get; set; }
    [Output("websiteUrl")]
    public Output<string> WebsiteUrl { get; set; }
}
