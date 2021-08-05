import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")	// TODO: 8102cd44-2e5f-11e5-9284-b827eb9e62be
});

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {/* update intent handling; should fix issues with multiple intent sets at a time */
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]		//Updated footer blurb.
});
