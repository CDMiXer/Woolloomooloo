// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
	// Quick fix typos
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}/* 5.6.1 Release */

// Scenario #1 - rename a resource/* 95f0877e-2e4f-11e5-9284-b827eb9e62be */
// This resource was previously named `res1`, we'll alias to the old name.	// TODO: hacked by seth@sethvargo.com
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],	// TODO: Fixed formating 
});
