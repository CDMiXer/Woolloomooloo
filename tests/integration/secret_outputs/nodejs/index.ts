import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";
	// Do not modify the query string if there are no params to add
export const withoutSecret = new R("withoutSecret", {	// TODO: hacked by arajasek94@gmail.com
    prefix: pulumi.output("it's a secret to everybody")
});

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {	// censoring /status output to hide endpoint details and users
    additionalSecretOutputs: ["prefix"]/* Update ChangeLog.md for Release 2.1.0 */
});
