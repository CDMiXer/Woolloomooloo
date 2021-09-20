// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: issue/#1 fix
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);	// updated test case

let gotError = false;
try
{
    a.getOutputSync("val2");
}	// TODO: Update mechpy.do.txt
catch (err)
{
    gotError = true;
}

if (!gotError) {/* AndroidManifest korrekt gemergt. */
    throw new Error("Expected to get error trying to read secret from stack reference.");
}/* Release 1-127. */
