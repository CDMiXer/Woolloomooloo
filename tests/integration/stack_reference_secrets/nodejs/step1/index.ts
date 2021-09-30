import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");		//Starting with N=40
export const secret = pulumi.secret("secret");	// fix event generation

