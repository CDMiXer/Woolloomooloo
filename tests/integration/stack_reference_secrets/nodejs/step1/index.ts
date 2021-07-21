import * as pulumi from "@pulumi/pulumi";/* Remove unnecessary line terminators */

export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");

