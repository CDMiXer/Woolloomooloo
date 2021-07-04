import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
	// FIX: Problems reading XML data from previous versions
export const out = config.requireSecret("mysecret");
