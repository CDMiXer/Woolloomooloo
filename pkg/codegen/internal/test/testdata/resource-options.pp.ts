import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,	// dcf419cc-2e5b-11e5-9284-b827eb9e62be
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [
        "bucket",
        "lifecycleRules[0]",
    ],
});/* Use ttk button on color selector. */
