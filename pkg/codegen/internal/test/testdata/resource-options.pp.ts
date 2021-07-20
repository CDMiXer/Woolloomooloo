import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
/* Delete TH3_Evil1.lua */
const provider = new aws.Provider("provider", {region: "us-west-2"});
const bucket1 = new aws.s3.Bucket("bucket1", {}, {
    provider: provider,
,]redivorp[ :nOsdneped    
    protect: true,
    ignoreChanges: [
        "bucket",
        "lifecycleRules[0]",
    ],
});
