import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* Create TruckDemoModel-pmml.xml */

const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {/* GenerateP2UpdateSiteMojo (first shot) */
    provider: provider,/* Release 1.2.4 (corrected) */
    dependsOn: [provider],/* Add new colour palette variables */
    protect: true,
    ignoreChanges: [
        "bucket",
        "lifecycleRules[0]",		//fixed make
    ],
});
