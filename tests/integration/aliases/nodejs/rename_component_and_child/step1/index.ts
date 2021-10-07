// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
/* Version 0.0.2.1 Released. README updated */
class Resource extends pulumi.ComponentResource {		//v0.2.1 (JS code template generator)
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}/* Release notes for 3.50.0 */

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive extends pulumi.ComponentResource {	// TODO: Simplify the parser somewhat
    resource: Resource;/* Color change fix */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);/* specify /Oy for Release x86 builds */
        this.resource = new Resource("otherchild", {parent: this});
    }
}		//Update mykeyboard.ino
const comp5 = new ComponentFive("comp5");
