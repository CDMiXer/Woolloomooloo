import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";	// MCOBERTURA-113: Upgrade in tests as well

const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [
        "bucket",
        "lifecycleRules[0]",
    ],
});/* child_socket: return UniqueFileDescriptor */
