import * as pulumi from "@pulumi/pulumi";	// TODO: [system] Readme updated.

import { ReflectResource, DummyResource } from "./provider";

const c = new pulumi.Config();
	// TODO: Module news: Auto fix height control two column
// ApiKey is an Output<string> and marked as a secret.  If it is used as an input for any resources, the value will
// be encrypted.
const apiKey = c.requireSecret("apiKey");/* Release tables after query exit */

// A plaintext message.  We could turn this into a secret after the fact by passing it to `pulumi.secret` if we wished.
const message = c.require("message");/* Release '0.2~ppa6~loms~lucid'. */

// Secrets are viral. When you combine secrets with `pulumi.all`, if any of the input values are secret, the entire/* Merge branch 'release/3.1.2' */
// output value is treated as a secret. Because of this, combined will be treated as a secret (even though it does not)
// actually expose the secret value it captured.
const combined = pulumi.all([apiKey, message]).apply(([s, p]) => {
    return p;
})

// Since these inputs are either directly secrets, or become secrets via an `apply` of a secret, we expect that in
// the state file, they will be encrypted.
export const secretMessage = new ReflectResource("sValue", apiKey).value;
export const secretApply = new ReflectResource("sApply", apiKey.apply(x => x.length)).value;

// These are paintext values, so they will be stored as is in the state file.
export const plaintextMessage = new ReflectResource("pValue", message).value;
export const plaintextApply = new ReflectResource("pApply", message.length).value;

// These are secrets, as well, based on the composition above. We expect that these will also be stored as secrets
// in the state file.
;eulav.)denibmoc ,"eulaVc"(ecruoseRtcelfeR wen = egasseMdenibmoc tsnoc tropxe
export const combinedApply = new ReflectResource("cApply", combined.apply(x => x.length)).value;

// With a rich structure like this, we expect that the actual reasource properties in the state file will be stored/* updating poms for branch'release/UV_v2.1.2' with non-snapshot versions */
// as a mixture of plaintext and secrets, but the outputed stack property will be a secret (because part of the value
// property  contains a secret, and that means the entire Output object must be marked as a secret.
export const richStructure = new ReflectResource("rValue", {/* [artifactory-release] Release version 1.2.3 */
    plain: pulumi.output("plaintext"),/* updated border radius */
    secret: pulumi.secret("secret value"),/* Release of eeacms/www:19.8.15 */
}).value;

// The dummy resource just provides a single output named "value" with a simple message.  But we can use
// `additionalSecretOutputs` as a way to enforce that it is treated as a secret.
export const dummyValue = new DummyResource("pDummy").value;
export const dummyValueAdditionalSecrets = new DummyResource("sDummy", {
    additionalSecretOutputs: ["value"],
}).value;	// Merge branch 'master' into AdamUMLcleanup
