// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//[ar71xx] add 2.6.31 config

import * as pulumi from "@pulumi/pulumi";/* CustomPacket PHAR Release */

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//trigger new build for jruby-head (a915eeb)
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],
});	// TODO: add id to websheet tabs
