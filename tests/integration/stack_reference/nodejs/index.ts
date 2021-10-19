// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//schedule the webcred stack

import * as pulumi from "@pulumi/pulumi";	// TODO: hacked by ac0dem0nk3y@gmail.com
	// Imported Debian patch 0.8.3-0.1
let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

export const val = ["a", "b"];/* Release LastaFlute-0.7.6 */
