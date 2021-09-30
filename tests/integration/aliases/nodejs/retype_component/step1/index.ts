// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {	// Allow using non-table alias as a rhs relation name, fix for #84 and #59
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);		//[PSDK] Update wincodec.idl. CORE-11368
    }
}	// mouse wheel scrolling, page up, page down

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {/* Release 0.8.5. */
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}	// TODO: will be fixed by steven@stebalien.com
const comp4 = new ComponentFour("comp4");
