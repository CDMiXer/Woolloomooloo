// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Sticky header example */
import * as pulumi from "@pulumi/pulumi";
		//Started the regex literal.
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}
	// TODO: hacked by steven@stebalien.com
// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],
});
