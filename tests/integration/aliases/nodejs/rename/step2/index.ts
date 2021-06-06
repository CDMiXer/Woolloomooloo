// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
	// Update xpath
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}	// TODO: Added Apache library
/* Release 4.2.0 */
// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.	// TODO: hacked by 13860583249@yeah.net
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],
});		//Create MissingInteger.rb
