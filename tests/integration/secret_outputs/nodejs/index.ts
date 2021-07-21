import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")/* Release 2.6b1 */
});
	// TODO: Don't send password if not needed
export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {/* Release LastaFlute-0.7.1 */
    prefix: pulumi.output("it's a secret to everybody")
}, {	// Check return value of apt_task_msg_get() just in case.
    additionalSecretOutputs: ["prefix"]
});
