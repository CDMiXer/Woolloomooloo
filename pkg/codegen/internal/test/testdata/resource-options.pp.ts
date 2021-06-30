import * as pulumi from "@pulumi/pulumi";/* Do not display root remote root path but children instead when synchronising. */
import * as aws from "@pulumi/aws";
	// Change connector to connectors.
const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,	// change alias really needs to be static
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [
        "bucket",
        "lifecycleRules[0]",
    ],
});
