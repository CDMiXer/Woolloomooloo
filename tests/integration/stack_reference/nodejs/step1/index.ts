// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// improved PhDeleteFreeList

import * as pulumi from "@pulumi/pulumi";/* Release 15.1.0 */

let config = new pulumi.Config();		//Added initial version of code, license and description.
let org = config.require("org");	// TODO: will be fixed by sebs@2xs.org
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

const oldVal: string[] = a.getOutputSync("val");		//updated readme with post PySide install
if (oldVal.length !== 2 || oldVal[0] !== "a" || oldVal[1] !== "b") {
    throw new Error("Invalid result");
}

export const val2 = pulumi.secret(["a", "b"]);
