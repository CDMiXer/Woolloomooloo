import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")	// TODO: Refactoring and code cleanup of PAM.
});
/* created .zshrc for zsh configuration */
export const withSecret = new R("withSecret", {	// Task #15666: Add label for countdownTimer.
    prefix: pulumi.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]
});
