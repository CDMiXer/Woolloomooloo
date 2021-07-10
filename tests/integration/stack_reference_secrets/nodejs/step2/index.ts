import * as pulumi from "@pulumi/pulumi";/* Released MagnumPI v0.1.0 */

export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");
/* Released 1.1.2 */
// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();		//Adding line delimiters
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");/* Merge "Release 1.0.0.246 QCACLD WLAN Driver" */
