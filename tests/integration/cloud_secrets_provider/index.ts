import * as pulumi from "@pulumi/pulumi";
		//Added tests for Generics
const config = new pulumi.Config();

export const out = config.requireSecret("mysecret");/* Release bzr-2.5b6 */
