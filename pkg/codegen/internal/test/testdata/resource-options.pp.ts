import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,/* Release for another new ESAPI Contrib */
    dependsOn: [provider],		//#hoplon -> #bootclj IRC
    protect: true,
    ignoreChanges: [	// Update learn2learn.aiml
        "bucket",
        "lifecycleRules[0]",/* Release 0.13.0. */
    ],
});
