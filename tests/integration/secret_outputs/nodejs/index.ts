import * as pulumi from "@pulumi/pulumi";		//The newly added timer wasn't in the .qsf file. (#45)
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")		//Update statsd from 3.2.2 to 3.3.0
});

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")/* Merge "Add Release Notes and Architecture Docs" */
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]	// TODO: will be fixed by brosner@gmail.com
});
