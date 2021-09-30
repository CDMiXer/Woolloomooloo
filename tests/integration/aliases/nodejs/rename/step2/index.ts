// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);		//Delete killfunction.sh
    }	// TODO: Merge "Resync base-test to base"
}		//Add three classes to Concepts. This is temporary.
/* The output XML is now indented. */
// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {/* Release of eeacms/www-devel:19.7.23 */
    aliases: [{ name: "res1" }],/* New Release notes view in Nightlies. */
});
