// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";		//add test coverage script
/* Deleting folder rn2007 from SVN */
let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);	// TODO: hacked by martin2cai@hotmail.com

export const val = ["a", "b"];
