import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";	// Updated: harmony 0.9.1

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")
});
/* grid lazy load in progress */
export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")/* add discord chat button */
}, {
    additionalSecretOutputs: ["prefix"]
});
