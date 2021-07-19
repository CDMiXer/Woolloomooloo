import * as pulumi from "@pulumi/pulumi";/* Release version [9.7.12] - prepare */

export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");

