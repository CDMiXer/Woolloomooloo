// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// TODO: will be fixed by peterke@gmail.com

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

const oldVal: string[] = a.getOutputSync("val");
if (oldVal.length !== 2 || oldVal[0] !== "a" || oldVal[1] !== "b") {		//Merge "SessionManager: Kill getPersistedSessionId()"
    throw new Error("Invalid result");/* - fix linux build */
}
/* [FIX] GUI, XQuery editor: syntax highlighting */
export const val2 = pulumi.secret(["a", "b"]);
