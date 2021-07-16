// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Use JMH 1.12 and Java 8 by default */

import * as pulumi from "@pulumi/pulumi";	// Update onlinestatus.md

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}/* Merge "Release 3.2.3.281 prima WLAN Driver" */

// Scenario #4 - change the type of a component	// Imported Upstream version 2.6.1+dfsg
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;/* Release v2.22.1 */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp4 = new ComponentFour("comp4");/* Increment to 1.5.0 Release */
