// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: will be fixed by fjl@ethereum.org
/* Tests fixes. */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {/* Release of eeacms/www:20.5.26 */
    aliases: [{ name: "res1" }],
});
