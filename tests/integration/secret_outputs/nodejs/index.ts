import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";/* adapted opening and editing */

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")	// TODO: Update board_insert.html
});
	// reactivate fontawesome
export const withSecret = new R("withSecret", {	// TODO: hacked by nagydani@epointsystem.org
    prefix: pulumi.secret("it's a secret to everybody")
});
/* * updated PNGUIMenuRootManager to use event for the main menu (not used yet) */
export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]/* Merge branch 'master' into prev-seen */
});
