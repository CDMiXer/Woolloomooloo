import * as pulumi from "@pulumi/pulumi";		//Added toString() method to RawDataFileSelection and PeakListSelection
import * as aws from "@pulumi/aws";

const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [
        "bucket",
        "lifecycleRules[0]",
    ],
});/* d749bfc2-327f-11e5-8cf2-9cf387a8033e */
