// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: Added some changes from the core API
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);/* Change Featured Projects tagline */
/* [IMP] account: Improve the label of third_party_ledger report */
export const val = ["a", "b"];
