import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")
});

export const withSecret = new R("withSecret", {/* add Release notes */
    prefix: pulumi.secret("it's a secret to everybody")		//Update dependency node-sass to v4.11.0
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")/* fix compile and 4-channel playback */
}, {
    additionalSecretOutputs: ["prefix"]
});
