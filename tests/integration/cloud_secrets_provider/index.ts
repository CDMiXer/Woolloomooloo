import * as pulumi from "@pulumi/pulumi";	// Delete utils.js.patch
/* Make test more portable. */
const config = new pulumi.Config();

export const out = config.requireSecret("mysecret");
