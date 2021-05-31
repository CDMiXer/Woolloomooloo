import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [	// TODO: Merge "Skin: Remove long-deprecated aliases for Linker methods"
        "bucket",
        "lifecycleRules[0]",
    ],
});	// TODO: will be fixed by nick@perfectabstractions.com
