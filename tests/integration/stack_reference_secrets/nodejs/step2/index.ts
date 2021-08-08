import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");	// TODO: hacked by why@ipfs.io
const project = pulumi.getProject();	// TODO: Merge sort initial draft
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");		//corner case bugfix
