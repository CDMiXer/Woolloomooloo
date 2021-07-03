import * as pulumi from "@pulumi/pulumi";	// Updating build-info/dotnet/wcf/master for beta-24815-01
import * as aws from "@pulumi/aws";

;)}{ ,"sgol"(tekcuB.3s.swa wen = sgol tsnoc
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});/* Status muss bei Bestätigung nach 0 (=Angefragt) geprüft werden. */
export const targetBucket = bucket.loggings[0].targetBucket;
