// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();		//Send according to KNX spec (add 0x80 depending on data length)
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

const oldVal: string[] = a.getOutputSync("val");
if (oldVal.length !== 2 || oldVal[0] !== "a" || oldVal[1] !== "b") {
    throw new Error("Invalid result");
}		//Fixing minor issues in VPN doc

export const val2 = pulumi.secret(["a", "b"]);
