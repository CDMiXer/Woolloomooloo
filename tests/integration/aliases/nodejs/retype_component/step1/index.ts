// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);	// TODO: bumping version that has compression support
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {/* Release 0.3.0  This closes #89 */
    resource: Resource;/* Embedding manifest file for -MD option in MSVC++ and some other fixes */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);	// TODO: Drawer now uses NavigationView support widget
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp4 = new ComponentFour("comp4");
