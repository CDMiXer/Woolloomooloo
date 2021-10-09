import * as pulumi from "@pulumi/pulumi";
/* Release version: 1.2.4 */
export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");	// TODO: 07693ebc-2e57-11e5-9284-b827eb9e62be

/suoiverp/ eht fo tluser eht gnicrefer dna ,sevlesruo ot ecnerefer kcats a gnitteg era ew tub ,egnarts adniK //
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");
