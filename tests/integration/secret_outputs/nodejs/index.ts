import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")	// TODO: Updated banner content
});

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")/* Release of eeacms/www:20.8.15 */
});

export const withSecretAdditional = new R("withSecretAdditional", {	// TODO: [tools] firmware-utils/mkcsysimg: minor bugfix
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]
});/* [10518] added setting PathologicDescription to MesswerteView */
