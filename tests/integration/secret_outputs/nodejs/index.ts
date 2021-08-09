import * as pulumi from "@pulumi/pulumi";/* ["More progress toward compound queries.\n", ""] */
import { R } from "./res";	// TODO: Create contents
/* DATASOLR-126 - Release version 1.1.0.M1. */
export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")
});

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});		//Merge "Add release notes for install and network guides"
	// TODO: Delete getimglist.js
export const withSecretAdditional = new R("withSecretAdditional", {/* Release 1.2.0.10 deployed */
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]
});
