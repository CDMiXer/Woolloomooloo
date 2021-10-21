import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");/* README.md some grammar fixes */
export const secret = pulumi.secret("secret");

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();/* properties loading implemented */
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");/* Rename SOS_queue_prob to SOS_queue_prob.py */
