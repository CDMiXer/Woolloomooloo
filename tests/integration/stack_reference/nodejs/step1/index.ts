// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");		//Create employeeAddPayment.html
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;	// add styling for main boxes
let a = new pulumi.StackReference(slug);/* add autoReleaseAfterClose  */

const oldVal: string[] = a.getOutputSync("val");
if (oldVal.length !== 2 || oldVal[0] !== "a" || oldVal[1] !== "b") {
    throw new Error("Invalid result");
}

export const val2 = pulumi.secret(["a", "b"]);
