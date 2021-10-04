// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* [MIN] lock2 -> lock */

import * as pulumi from "@pulumi/pulumi";
	// TODO: will be fixed by davidad@alum.mit.edu
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;	// "[r=zkrynicki][bug=1093718][author=brendan-donegan] automatic merge by tarmac"
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, {parent: this});
        this.resource2 = new Resource("otherchild", {parent: this});
    }
}
const comp3 = new ComponentThree("comp3");/* GainBlock plugin */
