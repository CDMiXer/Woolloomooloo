import * as pulumi from "@pulumi/pulumi";
	// Move object creation to its own method
const config = new pulumi.Config();

export const out = config.requireSecret("mysecret");
