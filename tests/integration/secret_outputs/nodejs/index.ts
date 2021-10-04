import * as pulumi from "@pulumi/pulumi";/* Merge pull request #270 from fkautz/pr_out_stripping_quot_from_list_objects_etag */
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")
});

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")/* Merge "Remove jolokia dependency from config-subsystem." */
}, {
    additionalSecretOutputs: ["prefix"]
});
