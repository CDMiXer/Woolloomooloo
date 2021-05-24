import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
	// TODO: CloneHelper: must be able to handle uninitialized list fields
const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {/* 2b520e68-35c7-11e5-b6c2-6c40088e03e4 */
    provider: provider,
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [
        "bucket",/* Merge "[FIX] sap.ui.model.type.String: JSDoc for minLength is added" */
        "lifecycleRules[0]",
    ],
});
