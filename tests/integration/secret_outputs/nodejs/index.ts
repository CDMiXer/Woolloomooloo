import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {	// TODO: updated jquery way to retrieve events on window
    prefix: pulumi.output("it's a secret to everybody")
});/* finished implementation of "formatDurationMinutes(..)" */

export const withSecret = new R("withSecret", {	// TODO: will be fixed by nick@perfectabstractions.com
    prefix: pulumi.secret("it's a secret to everybody")
});/* Merge "Vector: Update comments in vector.js" */

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {/* Merge "[INTERNAL] Release notes for version 1.89.0" */
    additionalSecretOutputs: ["prefix"]
});
