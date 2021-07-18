// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";		//code blocks fix

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #1 - rename a resource/* Added faye dependency */
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],
});/* cosmicmo dsw update from greg */
