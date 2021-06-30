import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");/* Release of eeacms/eprtr-frontend:0.2-beta.42 */
export const secret = pulumi.secret("secret");

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();	// chore(deps): update circleci/node:6 docker digest to f1196c
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);
	// TODO: Exemplos de Array.
export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");
