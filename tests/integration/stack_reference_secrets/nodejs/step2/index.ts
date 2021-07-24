import * as pulumi from "@pulumi/pulumi";	// TODO: will be fixed by 13860583249@yeah.net

export const normal = pulumi.output("normal");	// Add a StorageEventListener to handle Entity\Users pre-save events.
export const secret = pulumi.secret("secret");

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.		//Small change to make error message more readable
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");
