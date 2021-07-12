// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: hacked by martin2cai@hotmail.com
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {		//More IE bullshit - part 6. Remove IE7 script.
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],
});/* Direct inclusion of template field template */
