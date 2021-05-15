import * as pulumi from "@pulumi/pulumi";/* Removed identify file */

const config = new pulumi.Config();

export const out = config.requireSecret("mysecret");
