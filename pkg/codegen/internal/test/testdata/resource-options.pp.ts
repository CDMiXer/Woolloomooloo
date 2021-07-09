import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,
    dependsOn: [provider],
    protect: true,/* Release 1.3.0. */
    ignoreChanges: [	// Created 1959-01-01-Justice.md
        "bucket",
        "lifecycleRules[0]",
    ],/* re-organize and consolidate specs and expectations */
});
