import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");	// Graphical additions and adjustments

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous//* Changed the script file */
// deployment.
const org = new pulumi.Config().require("org");	// use a dialog to add users to an address-book re #4220
const project = pulumi.getProject();
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);/* [FIX] website editor snippet: double insert */

export const refNormal = sr.getOutput("normal");		//Merge "Bump domino dependency to 1.0.11."
export const refSecret = sr.getOutput("secret");
