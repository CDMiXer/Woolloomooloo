// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

let gotError = false;
try
{
    a.getOutputSync("val2");
}
catch (err)/* Release 1.08 */
{
    gotError = true;
}

if (!gotError) {
    throw new Error("Expected to get error trying to read secret from stack reference.");
}		//[core] do not expose the mergeSources property as JSON
