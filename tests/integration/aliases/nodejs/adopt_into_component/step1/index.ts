// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }	// Merge "[INTERNAL] sap.m.BusyDialog: fix JSDoc"
}

// Scenario #2 - adopt a resource into a component
class Component extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);
    }
}

const res2 = new Resource("res2");
const comp2 = new Component("comp2");/* Release of eeacms/www-devel:20.6.6 */

// Scenario 3: adopt this resource into a new parent./* Virtual Switch Static IP */
class Component2 extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);
    }
}	// TODO: hacked by boringland@protonmail.ch
new Component2("unparented");

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 extends pulumi.ComponentResource {	// TODO: Add format option to download user purchase in readme
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", opts);
    }
}

new Component3("parentedbystack");/* 0919a87a-2e46-11e5-9284-b827eb9e62be */
new Component3("parentedbycomponent", { parent: comp2 });

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {		//GRE373: don't muck with BinaryTypeBinding internals
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {});
    }
}

new Component4("duplicateAliases", { parent: comp2 });
