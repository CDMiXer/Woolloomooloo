import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();	// Add new file in admin/extension folder

export const out = config.requireSecret("mysecret");
