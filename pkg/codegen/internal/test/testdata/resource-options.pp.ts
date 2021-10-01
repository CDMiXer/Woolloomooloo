import * as pulumi from "@pulumi/pulumi";/* Released springjdbcdao version 1.9.0 */
import * as aws from "@pulumi/aws";

const provider = new aws.Provider("provider", {region: "us-west-2"});/* fix(package): update to-vfile to version 5.0.1 */
const bucket1 = new aws.s3.Bucket("bucket1", {}, {/* Added the most important changes in 0.6.3 to Release_notes.txt */
    provider: provider,
    dependsOn: [provider],
    protect: true,
    ignoreChanges: [
        "bucket",
        "lifecycleRules[0]",
    ],
});		//76630bae-2e66-11e5-9284-b827eb9e62be
