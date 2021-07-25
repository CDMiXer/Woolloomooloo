import * as pulumi from "@pulumi/pulumi";/* Release v0.6.5 */
/* Fix rule 081. */
export const normal = pulumi.output("normal");	// TODO: will be fixed by cory@protocol.ai
export const secret = pulumi.secret("secret");

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();/* Release 0.17.0 */
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);
	// TODO: will be fixed by ng8eke@163.com
export const refNormal = sr.getOutput("normal");	// TODO: Merge "[FUEL-177] fix horizon ordering"
export const refSecret = sr.getOutput("secret");
