// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* ProRelease2 update R11 should be 470 Ohm */

class Resource extends pulumi.ComponentResource {/* Add Release History to README */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: 13b59b54-2e6e-11e5-9284-b827eb9e62be
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {		//Added button icons
    aliases: [{ name: "res1" }],
});	// TODO: hacked by sebs@2xs.org
