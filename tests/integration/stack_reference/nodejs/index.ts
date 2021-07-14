// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: hacked by mowrain@yandex.com

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");	// TODO: working on dependency injection
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

export const val = ["a", "b"];
