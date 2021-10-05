// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release version [10.6.3] - prepare */
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

export const val = ["a", "b"];/* change db connect to h2 */
