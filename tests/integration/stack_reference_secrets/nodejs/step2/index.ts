import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");/* update chat demo */
export const secret = pulumi.secret("secret");
		//Fix generation of CHM name for release candidates.
// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);
		//merged [19626] to UOS 2.1
export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");	// TODO: will be fixed by remco@dutchcoders.io
