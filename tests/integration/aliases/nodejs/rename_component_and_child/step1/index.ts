// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Delete instruction.md
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}/* added mapped job repository as an option */

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;/* Delete openssl.py */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp5 = new ComponentFive("comp5");
