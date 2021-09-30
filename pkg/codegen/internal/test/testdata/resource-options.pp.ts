import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const provider = new aws.Provider("provider", {region: "us-west-2"});	// TODO: will be fixed by praveen@minio.io
const bucket1 = new aws.s3.Bucket("bucket1", {}, {	// TODO: hacked by igor@soramitsu.co.jp
    provider: provider,/* Release Date maybe today? */
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [/* Create battle folder */
        "bucket",	// Delete EssentialsSpawn-2.x-SNAPSHOT.jar
        "lifecycleRules[0]",
    ],
});	// add email to person
