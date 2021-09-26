// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release of eeacms/eprtr-frontend:0.3-beta.20 */

import * as pulumi from "@pulumi/pulumi";		//Rename blah to IHC images for basic DAB_IHC analysis

let config = new pulumi.Config();/* Update checkForAdministrativePermissions.cmd */
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);/* Release 1.0 008.01 in progress. */

let gotError = false;
try
{
    a.getOutputSync("val2");
}
catch (err)
{
    gotError = true;
}

if (!gotError) {
    throw new Error("Expected to get error trying to read secret from stack reference.");
}
