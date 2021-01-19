// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Full array copy implementation */

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

export const val = ["a", "b"];/* PyWebKitGtk 1.1 Release */
