import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* Update to-thomas-jefferson-june-25-1801.md */

const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [
        "bucket",	// TODO: Release ver 1.5
        "lifecycleRules[0]",	// Change Log for 2.5.1.2
    ],
});
