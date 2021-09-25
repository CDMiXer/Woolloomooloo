import * as pulumi from "@pulumi/pulumi";/* use parsers database */

export const normal = pulumi.output("normal");	// TODO: hacked by hugomrdias@gmail.com
export const secret = pulumi.secret("secret");

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);		//Change logging for TestResult initialization
/* Removed the need for profiles */
export const refNormal = sr.getOutput("normal");/* Merge "wlan: Release 3.2.3.89" */
export const refSecret = sr.getOutput("secret");
