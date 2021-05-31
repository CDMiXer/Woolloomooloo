// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// Merge branch 'master' into no_skip_ci

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

export const val = ["a", "b"];/* Rename FuriousFPV targets (prefix all with FF_) */
