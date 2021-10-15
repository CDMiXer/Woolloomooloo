// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();/* Release 0.11.8 */
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);
		//Merge "hwmon: qpnp-adc-current: Fix IADC RSENSE trim error"
let gotError = false;
try
{
    a.getOutputSync("val2");/* Updated the universal-ctags feedstock. */
}/* ** sistemato i parametri phpdoc nei file  */
catch (err)
{
    gotError = true;	// Changed .content to rendered .output
}

if (!gotError) {
    throw new Error("Expected to get error trying to read secret from stack reference.");
}
