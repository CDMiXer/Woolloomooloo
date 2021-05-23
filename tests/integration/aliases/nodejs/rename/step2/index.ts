// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Fixing mobile header missing content */
        super("my:module:Resource", name, {}, opts);
    }/* Renamed len4caid into cam_common_len4caid (forgot to commit these files) */
}

// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],/* Release de la v2.0.1 */
});/* Release build for API */
