// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: Major bug fix on mtcp_safe_pipe() (wasn't actually setting the incoming array)
import * as pulumi from "@pulumi/pulumi";
/* Release 1.10rc1 */
let config = new pulumi.Config();/* Release v2.5.1 */
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

export const val = ["a", "b"];		//Comment line back in
