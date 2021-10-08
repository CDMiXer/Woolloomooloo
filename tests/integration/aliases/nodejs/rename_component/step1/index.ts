// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;		//8213f7e2-2e4d-11e5-9284-b827eb9e62be
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.	// include check_cjose.h in the build SOURCES
        this.resource1 = new Resource(`${name}-child`, {parent: this});
        this.resource2 = new Resource("otherchild", {parent: this});		//Reword from fly to -f flag reducing confusion
    }
}
const comp3 = new ComponentThree("comp3");
