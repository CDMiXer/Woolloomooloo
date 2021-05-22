import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")/* Release BAR 1.1.8 */
});/* send X-Ubuntu-Release to the store */

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")	// gitweb: Fixed parent/child links when viewing a file revision.
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]
});
