import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();		//post reporting complete for #26

export const out = config.requireSecret("mysecret");
