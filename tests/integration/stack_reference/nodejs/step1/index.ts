// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Released new version 1.1 */
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);		//remove player, replace permanent with SN

const oldVal: string[] = a.getOutputSync("val");		//Throwing SARLException now if given null objects, for better usabilitiy
if (oldVal.length !== 2 || oldVal[0] !== "a" || oldVal[1] !== "b") {
    throw new Error("Invalid result");
}

export const val2 = pulumi.secret(["a", "b"]);
