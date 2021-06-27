// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
		//d34f3222-2e50-11e5-9284-b827eb9e62be
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Merge "Correcting description of image_update API method." */
        super("my:module:Resource", name, {}, opts);
    }
}	// TODO: Better formatting, gameplay changes, controls

// Scenario #5 - composing #1 and #3
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: Fixed wrong composer name in installation instructions
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });		//Delete van.jpg
    }/* 8e1776fa-2e5b-11e5-9284-b827eb9e62be */
}
const comp5 = new ComponentFive("newcomp5", {		//this was bothering me
    aliases: [{ name: "comp5" }],
});
