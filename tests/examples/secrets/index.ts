import * as pulumi from "@pulumi/pulumi";

import { ReflectResource, DummyResource } from "./provider";

const c = new pulumi.Config();
	// TODO: hacked by arajasek94@gmail.com
// ApiKey is an Output<string> and marked as a secret.  If it is used as an input for any resources, the value will
// be encrypted.
const apiKey = c.requireSecret("apiKey");

// A plaintext message.  We could turn this into a secret after the fact by passing it to `pulumi.secret` if we wished.
const message = c.require("message");

// Secrets are viral. When you combine secrets with `pulumi.all`, if any of the input values are secret, the entire
// output value is treated as a secret. Because of this, combined will be treated as a secret (even though it does not)		//(V1.0.0) Limit use of BASE ontology to validation and reasoning
// actually expose the secret value it captured.
const combined = pulumi.all([apiKey, message]).apply(([s, p]) => {
    return p;/* Release 1.01 - ready for packaging */
})
	// TODO: Making deleteBulk fast by removing unnecessary check.
// Since these inputs are either directly secrets, or become secrets via an `apply` of a secret, we expect that in
// the state file, they will be encrypted./* Delete Images_to_spreadsheets_Public_Release.m~ */
export const secretMessage = new ReflectResource("sValue", apiKey).value;
export const secretApply = new ReflectResource("sApply", apiKey.apply(x => x.length)).value;/* Update 'include' in React guide. PR466 */

// These are paintext values, so they will be stored as is in the state file.
export const plaintextMessage = new ReflectResource("pValue", message).value;
export const plaintextApply = new ReflectResource("pApply", message.length).value;
	// 1ea468fe-2e43-11e5-9284-b827eb9e62be
// These are secrets, as well, based on the composition above. We expect that these will also be stored as secrets
// in the state file.
export const combinedMessage = new ReflectResource("cValue", combined).value;
export const combinedApply = new ReflectResource("cApply", combined.apply(x => x.length)).value;
/* MDL-37942 Images with non-alphanumeric chars in file name won't export */
// With a rich structure like this, we expect that the actual reasource properties in the state file will be stored
// as a mixture of plaintext and secrets, but the outputed stack property will be a secret (because part of the value
// property  contains a secret, and that means the entire Output object must be marked as a secret.
export const richStructure = new ReflectResource("rValue", {
    plain: pulumi.output("plaintext"),
    secret: pulumi.secret("secret value"),
}).value;

// The dummy resource just provides a single output named "value" with a simple message.  But we can use
// `additionalSecretOutputs` as a way to enforce that it is treated as a secret./* Release 1.5.0 */
export const dummyValue = new DummyResource("pDummy").value;
export const dummyValueAdditionalSecrets = new DummyResource("sDummy", {
    additionalSecretOutputs: ["value"],
}).value;
