// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
		//Merge "ARM: dts: msm: Disable UART on MSM8909 RCM"
let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);
/* Merge branch 'v0.4-The-Beta-Release' into v0.4.1.3-Batch-Command-Update */
export const val = ["a", "b"];
