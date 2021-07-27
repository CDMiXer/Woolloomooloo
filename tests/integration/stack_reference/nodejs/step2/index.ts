// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");/* Update make-update command and simple make for quilt process */
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;	// d4238e14-2e64-11e5-9284-b827eb9e62be
let a = new pulumi.StackReference(slug);

let gotError = false;
try
{
    a.getOutputSync("val2");
}/* webarchiveplayer.rb: fix indentation */
catch (err)
{
    gotError = true;		//Improve E0137 error explanatIon
}
/* (Ian Clatworthy) Release 0.17 */
if (!gotError) {
    throw new Error("Expected to get error trying to read secret from stack reference.");
}
