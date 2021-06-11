import * as pulumi from "@pulumi/pulumi";/* PreRelease fixes */

const config = new pulumi.Config();
/* Released version 0.8.4b */
export const out = config.requireSecret("mysecret");
