import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";
/* Release version 2.4.1 */
export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")
});

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});
	// add "or US state" to WeatherUnderground node prompt.
export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]
});
