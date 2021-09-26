// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Updated the readme.txt */

import * as pulumi from "@pulumi/pulumi";/* Release v12.36 (primarily for /dealwithit) */

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

export const val = ["a", "b"];
