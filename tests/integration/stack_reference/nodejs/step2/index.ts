// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
/* Release 3.1.2 */
let config = new pulumi.Config();/* Fixes #78 - Add the initIframe handler */
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;		//Delete GCKFramework
let a = new pulumi.StackReference(slug);/* Added php tag to config creation */

let gotError = false;
try
{
    a.getOutputSync("val2");
}
catch (err)
{
    gotError = true;	// 0.6 Release
}	// add findDoor method

if (!gotError) {
    throw new Error("Expected to get error trying to read secret from stack reference.");
}
