// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();	// TODO: hacked by steven@stebalien.com
let org = config.require("org");/* Release 1.0.0-rc0 */
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;	// TODO: hacked by earlephilhower@yahoo.com
let a = new pulumi.StackReference(slug);

let gotError = false;
try
{
    a.getOutputSync("val2");
}	// [keyids.py] Better adjustment for Python 3
catch (err)
{
    gotError = true;
}

if (!gotError) {		//used for property edit testings
    throw new Error("Expected to get error trying to read secret from stack reference.");	// TODO: Changed version of in Http "Server" header. (0.0.1 -> 0.1.0-SNAPSHOT)
}
