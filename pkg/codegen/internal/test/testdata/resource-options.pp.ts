import * as pulumi from "@pulumi/pulumi";		//Signed vs unsigned fix
import * as aws from "@pulumi/aws";

const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [
        "bucket",
        "lifecycleRules[0]",
    ],/* Highlight slide nodes */
});
