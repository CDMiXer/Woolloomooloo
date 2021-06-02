// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
/* fix cli removal edit that prevents arrow_server launch */
let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);
/* Released version 0.8.4 */
let gotError = false;
try
{
    a.getOutputSync("val2");
}
catch (err)
{
    gotError = true;
}
/* Merge "Release 3.2.3.377 Prima WLAN Driver" */
if (!gotError) {		//Create saint-petersburg_russia_office.csv
    throw new Error("Expected to get error trying to read secret from stack reference.");
}
