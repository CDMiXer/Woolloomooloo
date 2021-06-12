import * as pulumi from "@pulumi/pulumi";
/* Closes #144 */
export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");	// Remove the old grid layer

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();	// TODO: Fixed bugs(hopefully).
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");
