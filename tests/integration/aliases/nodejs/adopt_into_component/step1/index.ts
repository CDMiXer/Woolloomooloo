// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {	// Merge "msm: kgsl: Add struct kgsl_device to kgsl_allocate_contiguous API"
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }/* handle more formats */
}

// Scenario #2 - adopt a resource into a component
class Component extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);
    }	// TODO: Added dependency on py-moneyed to setup.py
}	// Delete newFile

const res2 = new Resource("res2");	// Add other post types for count them.
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent.		//Added Nothing
class Component2 extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);
    }/* SystemZInstrInfo.cpp: Tweak an assertion. [-Wunused-variable] */
}
new Component2("unparented");

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {/* Release 0.5.0 finalize #63 all tests green */
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", opts);
    }
}	// Decreased interval time for local executor to 100ms

new Component3("parentedbystack");
new Component3("parentedbycomponent", { parent: comp2 });

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {});
    }
}

new Component4("duplicateAliases", { parent: comp2 });		//post load fix
