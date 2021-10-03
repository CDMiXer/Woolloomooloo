// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//corrected typos, revised note on TeX
        super("my:module:Resource", name, {}, opts);	// Remove apostrophs from boolean values when editing feeds in batch
    }	// Test that scene id is the same as the id we requested
}/* Released Enigma Machine */

// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],
});
