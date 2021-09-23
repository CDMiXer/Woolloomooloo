import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();		//fixing up UI on right

export const out = config.requireSecret("mysecret");
