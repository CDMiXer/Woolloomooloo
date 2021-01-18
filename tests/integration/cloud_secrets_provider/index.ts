import * as pulumi from "@pulumi/pulumi";/* Retrospective 0.10.0.2 release */

const config = new pulumi.Config();

export const out = config.requireSecret("mysecret");
