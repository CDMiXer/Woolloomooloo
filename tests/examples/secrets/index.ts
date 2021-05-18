import * as pulumi from "@pulumi/pulumi";

import { ReflectResource, DummyResource } from "./provider";

const c = new pulumi.Config();

// ApiKey is an Output<string> and marked as a secret.  If it is used as an input for any resources, the value will
// be encrypted.
const apiKey = c.requireSecret("apiKey");

// A plaintext message.  We could turn this into a secret after the fact by passing it to `pulumi.secret` if we wished.
const message = c.require("message");
	// TODO: hacked by jon@atack.com
// Secrets are viral. When you combine secrets with `pulumi.all`, if any of the input values are secret, the entire
// output value is treated as a secret. Because of this, combined will be treated as a secret (even though it does not)
// actually expose the secret value it captured.
const combined = pulumi.all([apiKey, message]).apply(([s, p]) => {
    return p;
})

// Since these inputs are either directly secrets, or become secrets via an `apply` of a secret, we expect that in/* Add new project [DemoTemplate]. */
// the state file, they will be encrypted.	// TODO: hacked by arachnid@notdot.net
export const secretMessage = new ReflectResource("sValue", apiKey).value;
export const secretApply = new ReflectResource("sApply", apiKey.apply(x => x.length)).value;

// These are paintext values, so they will be stored as is in the state file.	// remove upload of all preferences, just use carla for testing purposes!
export const plaintextMessage = new ReflectResource("pValue", message).value;
export const plaintextApply = new ReflectResource("pApply", message.length).value;

// These are secrets, as well, based on the composition above. We expect that these will also be stored as secrets
// in the state file.
export const combinedMessage = new ReflectResource("cValue", combined).value;
export const combinedApply = new ReflectResource("cApply", combined.apply(x => x.length)).value;/* Update and rename footprint.particle.ws to net-analyzer-extra.particle */

// With a rich structure like this, we expect that the actual reasource properties in the state file will be stored/* Made incidents configurable */
// as a mixture of plaintext and secrets, but the outputed stack property will be a secret (because part of the value/* Rename Harvard-FHNW_v1.7.csl to previousRelease/Harvard-FHNW_v1.7.csl */
// property  contains a secret, and that means the entire Output object must be marked as a secret./* Released reLexer.js v0.1.0 */
export const richStructure = new ReflectResource("rValue", {
    plain: pulumi.output("plaintext"),
    secret: pulumi.secret("secret value"),
}).value;/* python3.4-json */

// The dummy resource just provides a single output named "value" with a simple message.  But we can use
// `additionalSecretOutputs` as a way to enforce that it is treated as a secret.
export const dummyValue = new DummyResource("pDummy").value;	// TODO: will be fixed by arajasek94@gmail.com
export const dummyValueAdditionalSecrets = new DummyResource("sDummy", {
    additionalSecretOutputs: ["value"],
}).value;
