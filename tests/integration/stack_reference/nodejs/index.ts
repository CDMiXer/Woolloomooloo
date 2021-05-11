// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// :palm_tree::rage: Updated in browser at strd6.github.io/editor
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");	// TODO: Fix spurious error in memory.limit (PR#13673), plus some cleanup of the docs
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;	// Typo fixed. Thank you @misabear
let a = new pulumi.StackReference(slug);

export const val = ["a", "b"];
