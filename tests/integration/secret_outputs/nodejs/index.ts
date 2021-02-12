import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")/* Automerge lp:~akopytov/percona-server/bug1248505 */
});
	// TODO: will be fixed by alan.shaw@protocol.ai
export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {	// TODO: 9dcff5ec-4b19-11e5-8758-6c40088e03e4
    additionalSecretOutputs: ["prefix"]
});
