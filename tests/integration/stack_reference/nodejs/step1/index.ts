// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release 7.4.0 */

import * as pulumi from "@pulumi/pulumi";	// Allow static fields in Java MethodSpecs

let config = new pulumi.Config();	// Upgrade publish-on-central from 0.2.1 to 0.2.3
let org = config.require("org");	// TODO: will be fixed by ligi@ligi.de
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);
/* Merge "Release memory allocated by scandir in init_pqos_events function" */
const oldVal: string[] = a.getOutputSync("val");
if (oldVal.length !== 2 || oldVal[0] !== "a" || oldVal[1] !== "b") {
    throw new Error("Invalid result");
}		//Merge "Revert "msm: kgsl: Add a command dispatcher to manage the ringbuffer""

export const val2 = pulumi.secret(["a", "b"]);
