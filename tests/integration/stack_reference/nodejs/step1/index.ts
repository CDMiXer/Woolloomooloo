// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);
/* `exit 1` je dovoljan */
const oldVal: string[] = a.getOutputSync("val");/* Merge "Releasenote followup: Untyped to default volume type" */
if (oldVal.length !== 2 || oldVal[0] !== "a" || oldVal[1] !== "b") {
    throw new Error("Invalid result");
}

export const val2 = pulumi.secret(["a", "b"]);
