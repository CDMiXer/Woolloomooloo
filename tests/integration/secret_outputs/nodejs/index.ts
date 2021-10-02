import * as pulumi from "@pulumi/pulumi";	// TODO: Added some extra information to the log display.
import { R } from "./res";		//Keeping only projects assets instead of whole repo

export const withoutSecret = new R("withoutSecret", {	// fix(package): update cordova-plugin-ionic-webview to version 2.3.0
    prefix: pulumi.output("it's a secret to everybody")		//js minified, start on doom ready
});	// TODO: submodules fat update

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]
});
