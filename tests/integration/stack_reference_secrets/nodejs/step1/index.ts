import * as pulumi from "@pulumi/pulumi";	// TODO: hacked by caojiaoyue@protonmail.com
	// TODO: Merge "Remove obvious function-level profiling"
export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");	// TODO: Add permission mobile update

