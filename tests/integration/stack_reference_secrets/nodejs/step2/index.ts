import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();/* Updated the Release notes with some minor grammar changes and clarifications. */
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");	// TODO: Added items to git/SVN ignore list [ci skip]
export const refSecret = sr.getOutput("secret");/* ui: make ship-mode=single-bibdata work again */
