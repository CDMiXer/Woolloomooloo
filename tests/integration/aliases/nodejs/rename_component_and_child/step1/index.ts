// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Adding round image */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {	// 27794e02-2e68-11e5-9284-b827eb9e62be
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time/* add infweb */
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp5 = new ComponentFive("comp5");
