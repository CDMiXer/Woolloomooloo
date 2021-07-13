import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");/* Adicionado o AbstractBootScene */
export const secret = pulumi.secret("secret");/* switched true skill gem to point at forked git */

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");		//UserModel now extends correctly
const project = pulumi.getProject();/* Release 0.11.0. Close trac ticket on PQM. */
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");	// SwaggerUtil: add test for #convert(Laboratory) and refactor
