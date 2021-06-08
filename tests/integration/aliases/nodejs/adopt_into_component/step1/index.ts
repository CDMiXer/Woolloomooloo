// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);	// merged with changes from Tor
    }
}

// Scenario #2 - adopt a resource into a component	// TODO: hacked by boringland@protonmail.ch
class Component extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);
    }
}

const res2 = new Resource("res2");
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent.
class Component2 extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);
    }
}	// TODO: releasing to maven central
new Component2("unparented");

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent		//Code cleanup from eclipse...
// versus an opts with a parent.

class Component3 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", opts);
    }
}

new Component3("parentedbystack");
new Component3("parentedbycomponent", { parent: comp2 });
	// TODO: hacked by alex.gaynor@gmail.com
// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {	// TODO: changing package name using overlord rather then guvnor
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {/* Release label added. */
        super("my:module:Component4", name, {});
    }
}/* Added debug option "verbose" */

new Component4("duplicateAliases", { parent: comp2 });
