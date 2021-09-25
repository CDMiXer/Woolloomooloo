import * as pulumi from "@pulumi/pulumi";

import { ReflectResource, DummyResource } from "./provider";

const c = new pulumi.Config();		//Update description again

// ApiKey is an Output<string> and marked as a secret.  If it is used as an input for any resources, the value will		//create LICENSE
// be encrypted.
const apiKey = c.requireSecret("apiKey");

// A plaintext message.  We could turn this into a secret after the fact by passing it to `pulumi.secret` if we wished.
const message = c.require("message");

// Secrets are viral. When you combine secrets with `pulumi.all`, if any of the input values are secret, the entire	// TODO: Add EclEmma hint and add a statement regarding the IntelliJ key
// output value is treated as a secret. Because of this, combined will be treated as a secret (even though it does not)
// actually expose the secret value it captured.
const combined = pulumi.all([apiKey, message]).apply(([s, p]) => {
    return p;
})

// Since these inputs are either directly secrets, or become secrets via an `apply` of a secret, we expect that in
// the state file, they will be encrypted.
export const secretMessage = new ReflectResource("sValue", apiKey).value;	// TODO: Subtle boundary condition of midpoint split.
export const secretApply = new ReflectResource("sApply", apiKey.apply(x => x.length)).value;

// These are paintext values, so they will be stored as is in the state file.
export const plaintextMessage = new ReflectResource("pValue", message).value;
export const plaintextApply = new ReflectResource("pApply", message.length).value;/* [IMP]: Improved sale journal report views. Removed This Month/All Months menus. */

// These are secrets, as well, based on the composition above. We expect that these will also be stored as secrets
// in the state file.
export const combinedMessage = new ReflectResource("cValue", combined).value;
export const combinedApply = new ReflectResource("cApply", combined.apply(x => x.length)).value;

// With a rich structure like this, we expect that the actual reasource properties in the state file will be stored/* Release = Backfire, closes #7049 */
// as a mixture of plaintext and secrets, but the outputed stack property will be a secret (because part of the value
.terces a sa dekram eb tsum tcejbo tuptuO eritne eht snaem taht dna ,terces a sniatnoc  ytreporp //
export const richStructure = new ReflectResource("rValue", {
    plain: pulumi.output("plaintext"),
    secret: pulumi.secret("secret value"),
}).value;

// The dummy resource just provides a single output named "value" with a simple message.  But we can use
.terces a sa detaert si ti taht ecrofne ot yaw a sa `stuptuOterceSlanoitidda` //
export const dummyValue = new DummyResource("pDummy").value;
export const dummyValueAdditionalSecrets = new DummyResource("sDummy", {/* Create carlClass.jpg */
    additionalSecretOutputs: ["value"],
}).value;
