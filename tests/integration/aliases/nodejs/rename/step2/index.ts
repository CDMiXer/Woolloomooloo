// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {/* Release notes for 1.0.62 */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.	// Delete Rules.md
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],
});
