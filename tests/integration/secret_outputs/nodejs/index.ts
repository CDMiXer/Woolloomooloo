import * as pulumi from "@pulumi/pulumi";/* v1.0.0 Release Candidate (today) */
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {/* List of algorithms added. */
    prefix: pulumi.output("it's a secret to everybody")
});
/* Release 0.94.200 */
export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});
		//Update the link to samples
export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")	// Create client_secrets.json
}, {
    additionalSecretOutputs: ["prefix"]
});
