// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Release v1.6.0 */
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;	// Fixed a bug that occured when final group of cycle only contains one cycle
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Add a missing [super dealloc]. */
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit/* Release 1.0.16 */
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, {parent: this});	// TODO: will be fixed by steven@stebalien.com
        this.resource2 = new Resource("otherchild", {parent: this});
    }
}
const comp3 = new ComponentThree("comp3");
