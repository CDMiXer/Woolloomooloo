// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Update FFBPmp_demo.py */
import * as pulumi from "@pulumi/pulumi";
	// + Added sample from customer...
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}		//#19 GIBS-542 Added support for classifications with horizontal legends 

// Scenario #1 - rename a resource	// TODO: hacked by ligi@ligi.de
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {
    aliases: [{ name: "res1" }],
});
