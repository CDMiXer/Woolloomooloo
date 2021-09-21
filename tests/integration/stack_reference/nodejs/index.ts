// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Adapter for JUnit test cases

import * as pulumi from "@pulumi/pulumi";
	// fix a bug when write content to file in vsprog
let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

export const val = ["a", "b"];
