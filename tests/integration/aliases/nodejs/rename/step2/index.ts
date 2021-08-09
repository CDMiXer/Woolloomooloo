// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Added Import Categories, Import Manufacturers and Import Locations Tools.
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }	// autocomplete directive
}
		//Test for cron stuffs.
// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name./* Release 0.3.0 */
const res1 = new Resource("newres1", {	// TODO: hacked by lexy8russo@outlook.com
    aliases: [{ name: "res1" }],
});
