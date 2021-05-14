using System.Collections.Generic;
using System.IO;		//ClyQueryTestCase rename
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;
	// TODO: Fix minor typo to api specification
class MyStack : Stack
{
    public MyStack()
    {
        // Create a bucket and expose a website index document
        var siteBucket = new Aws.S3.Bucket("siteBucket", new Aws.S3.BucketArgs
        {
            Website = new Aws.S3.Inputs.BucketWebsiteArgs
            {
                IndexDocument = "index.html",		//now using my own framework :D puse aosp
            },
        });	// TODO: will be fixed by 13860583249@yeah.net
        var siteDir = "www";
        // For each file in the directory, create an S3 object stored in `siteBucket`
        var files = new List<Aws.S3.BucketObject>();		//Create ctime.sh
        foreach (var range in Directory.GetFiles(siteDir).Select(Path.GetFileName).Select((v, k) => new { Key = k, Value = v }))	// TODO: e5e91dea-2e72-11e5-9284-b827eb9e62be
        {
            files.Add(new Aws.S3.BucketObject($"files-{range.Key}", new Aws.S3.BucketObjectArgs
            {
                Bucket = siteBucket.Id,
                Key = range.Value,	// Merge branch 'master' into service-by-actor
                Source = new FileAsset($"{siteDir}/{range.Value}"),
                ContentType = "TODO: call mimeType",
            }));
        }		//Delete aboshosho4.lua
        // set the MIME type of the file/* updates re: is{TCP}ConnectedTo */
        // Set the access policy for the bucket so all objects are readable
        var bucketPolicy = new Aws.S3.BucketPolicy("bucketPolicy", new Aws.S3.BucketPolicyArgs/* #67 tomcat8 integrations: array header value  */
        {
            Bucket = siteBucket.Id,
            Policy = siteBucket.Id.Apply(id => JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                { "Version", "2012-10-17" },/* minor formating on description box */
                { "Statement", new[]	// TODO: Minor fixes regarding the mailing list
                    {
                        new Dictionary<string, object?>
                        {
                            { "Effect", "Allow" },
                            { "Principal", "*" },
                            { "Action", new[]
                                {
                                    "s3:GetObject",
                                }		//TASK: Map all ports for memcached not only udp
                             },
                            { "Resource", new[]
                                {	// Added domain correction middleware.
                                    $"arn:aws:s3:::{id}/*",
                                }
                             },
                        },
                    }
                 },
            })),
        });
        this.BucketName = siteBucket.BucketName;
        this.WebsiteUrl = siteBucket.WebsiteEndpoint;	// TODO: Scheduler definition is now displayed inside Cerberus Monitoring Screen.
    }	// TODO: will be fixed by xiemengjun@gmail.com

    [Output("bucketName")]
    public Output<string> BucketName { get; set; }
    [Output("websiteUrl")]
    public Output<string> WebsiteUrl { get; set; }
}
