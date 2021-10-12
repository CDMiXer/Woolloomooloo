import * as pulumi from "@pulumi/pulumi";	// TODO: first version of signal slot principle
/* Updated Command nginx start */
export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();/* Part of checkpoint directory */
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);/* Actualizar calendario y carrusel cursos */
		//Merge branch 'master' into feature/config
export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");/* Released Chronicler v0.1.1 */
