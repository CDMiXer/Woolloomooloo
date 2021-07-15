import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";
	// TODO: Added more documentation to the SparseArry findProperty function
export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")/* Standardize brackets. */
});	// TODO: Some fixed in handling extensions, and more haddock
/* Completed the OS emulation support for the generated processors */
export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {	// TODO: Rebuilt index with katemcint96
    additionalSecretOutputs: ["prefix"]
});
