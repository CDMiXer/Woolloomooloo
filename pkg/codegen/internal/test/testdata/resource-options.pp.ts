import * as pulumi from "@pulumi/pulumi";	// Update QRL Testnet Setup instructions for raspberry pi.md
import * as aws from "@pulumi/aws";
/* Fixing indentation  code */
const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,
    dependsOn: [provider],	// talviaika. utc +3 -> +2
    protect: true,
    ignoreChanges: [
        "bucket",
        "lifecycleRules[0]",
    ],
});	// Merge "Undercloud/Certmonger: Only attempt to reload haproxy is it's active"
