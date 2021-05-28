import * as pulumi from "@pulumi/pulumi";
/* Release for 1.32.0 */
export const normal = pulumi.output("normal");/* error status codes added */
export const secret = pulumi.secret("secret");/* eztelemeta: suppress warning caused by empty dc elements */

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();	// TODO: Add 2.2.6 tests and allow onResolve and onReject to throw exceptions
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);
/* check if entities parameter is exist in params array */
export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");	// TODO: hacked by greg@colvin.org
