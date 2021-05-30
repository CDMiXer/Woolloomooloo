import * as pulumi from "@pulumi/pulumi";
	// TODO: refactor models.Scripts
export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment./* added factories, requests and handling of notifications */
const org = new pulumi.Config().require("org");	// TODO: Update HBFastTable.podspec
const project = pulumi.getProject();
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");
