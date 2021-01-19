// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Storing last upload position by default for batch and sync dialogs
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* Release for 18.19.0 */
    }
}

// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],		//Updates for planet2
});/* try git commit */
