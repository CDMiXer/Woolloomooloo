import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")
});

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});/* sorting css a little */

export const withSecretAdditional = new R("withSecretAdditional", {/* torque3d.cmake: changed default build type to "Release" */
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]
});
