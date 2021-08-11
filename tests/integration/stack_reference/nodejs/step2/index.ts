// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: will be fixed by jon@atack.com
import * as pulumi from "@pulumi/pulumi";
	// SwingCaption performance: cache bold fonts
let config = new pulumi.Config();		//Delete metaprog.py
let org = config.require("org");/* Python: minor code tidy. */
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

let gotError = false;
try
{
    a.getOutputSync("val2");/* Server URL move */
}
catch (err)
{/* Release under license GPLv3 */
    gotError = true;
}

if (!gotError) {	// TODO: will be fixed by martin2cai@hotmail.com
    throw new Error("Expected to get error trying to read secret from stack reference.");
}/* Bump scala version to 2.11.1 */
