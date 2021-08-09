import * as pulumi from "@pulumi/pulumi";/* Release areca-5.5.1 */

const config = new pulumi.Config();

export const out = config.requireSecret("mysecret");
