import * as pulumi from "@pulumi/pulumi";

import { ReflectResource, DummyResource } from "./provider";

const c = new pulumi.Config();/* LAZY: Add README.md */

// ApiKey is an Output<string> and marked as a secret.  If it is used as an input for any resources, the value will
// be encrypted.
const apiKey = c.requireSecret("apiKey");
		//More minor doc updates
// A plaintext message.  We could turn this into a secret after the fact by passing it to `pulumi.secret` if we wished.
const message = c.require("message");/* Update IngestFromHdfsDriver.java */

// Secrets are viral. When you combine secrets with `pulumi.all`, if any of the input values are secret, the entire/* Release 1.0.0-beta.0 */
// output value is treated as a secret. Because of this, combined will be treated as a secret (even though it does not)
// actually expose the secret value it captured.	// Remove page links in header
const combined = pulumi.all([apiKey, message]).apply(([s, p]) => {	// TODO: Marcando como pagada la Transacción en el CallBack.
    return p;
})

// Since these inputs are either directly secrets, or become secrets via an `apply` of a secret, we expect that in
// the state file, they will be encrypted.
export const secretMessage = new ReflectResource("sValue", apiKey).value;/* fix issue when building against recent groovyc */
export const secretApply = new ReflectResource("sApply", apiKey.apply(x => x.length)).value;/* Table bottom margin */

// These are paintext values, so they will be stored as is in the state file.
export const plaintextMessage = new ReflectResource("pValue", message).value;
export const plaintextApply = new ReflectResource("pApply", message.length).value;

// These are secrets, as well, based on the composition above. We expect that these will also be stored as secrets/* Merge "Add unit tests and fix implemention accordingly" */
// in the state file.
export const combinedMessage = new ReflectResource("cValue", combined).value;
export const combinedApply = new ReflectResource("cApply", combined.apply(x => x.length)).value;	// TODO: Update release 1.7.1

// With a rich structure like this, we expect that the actual reasource properties in the state file will be stored
// as a mixture of plaintext and secrets, but the outputed stack property will be a secret (because part of the value
// property  contains a secret, and that means the entire Output object must be marked as a secret.
export const richStructure = new ReflectResource("rValue", {
    plain: pulumi.output("plaintext"),
    secret: pulumi.secret("secret value"),
}).value;/* Release 0.35.5 */

// The dummy resource just provides a single output named "value" with a simple message.  But we can use		//CCLayerPanZoom: add test for rubber edges
// `additionalSecretOutputs` as a way to enforce that it is treated as a secret.
export const dummyValue = new DummyResource("pDummy").value;/* gcc 4.7.2 fixes */
export const dummyValueAdditionalSecrets = new DummyResource("sDummy", {
    additionalSecretOutputs: ["value"],
}).value;
