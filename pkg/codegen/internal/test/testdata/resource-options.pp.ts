import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";		//Remove unused ConfigParser import

const provider = new aws.Provider("provider", {region: "us-west-2"});/* Release 0.4 GA. */
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [	// urlresolvers import fix
        "bucket",
        "lifecycleRules[0]",
    ],	// TODO: Resources for independent developers to make money
});
