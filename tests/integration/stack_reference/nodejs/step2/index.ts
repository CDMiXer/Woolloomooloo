// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);
/* removed errant pdb */
let gotError = false;/* AbstractReturnValueFactory: adjustment */
try
{
    a.getOutputSync("val2");
}
catch (err)/* improving split fetch */
{
    gotError = true;
}

if (!gotError) {
    throw new Error("Expected to get error trying to read secret from stack reference.");	// Basic routing idea and gain check on datalogger
}
