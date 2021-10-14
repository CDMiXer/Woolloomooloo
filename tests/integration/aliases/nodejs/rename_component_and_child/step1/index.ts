// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: e390a200-2e42-11e5-9284-b827eb9e62be
        super("my:module:Resource", name, {}, opts);		//Create 423. Reconstruct Original Digits from English.java
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time/* Display Release build results */
class ComponentFive extends pulumi.ComponentResource {/* Added terminal ansi coloring as an option */
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}/* BI Fusion v3.0 Official Release */
const comp5 = new ComponentFive("comp5");
