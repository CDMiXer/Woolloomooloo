import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
	// TODO: hacked by mail@bitpshr.net
export const out = config.requireSecret("mysecret");
