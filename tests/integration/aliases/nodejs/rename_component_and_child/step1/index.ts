// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* JForum 2.3.3 Release */

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* Released springjdbcdao version 1.7.7 */
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);/* Released 0.2.2 */
        this.resource = new Resource("otherchild", {parent: this});
    }/* Release 1-90. */
}/* 0.9.7 Release. */
const comp5 = new ComponentFive("comp5");
