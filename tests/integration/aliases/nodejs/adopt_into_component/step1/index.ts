// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Removed useless user dir */

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #2 - adopt a resource into a component
class Component extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Delete daihoX.png */
        super("my:module:Component", name, {}, opts);
    }
}
/* Create QuotesList2Nacho */
const res2 = new Resource("res2");
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent./* Fix #1055129 (Xoom MTP 'Send to Main' send to SDCARD) */
class Component2 extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// fe98ec32-2e3e-11e5-9284-b827eb9e62be
        super("my:module:Component2", name, {}, opts);
    }
}
new Component2("unparented");/* add missing menu xml */
		//Add dialog when clear all records.
// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix/* switch to using bookmarklet passed in by view */
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent./* Added yahoo mail configuration */

class Component3 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);	// TODO: hacked by nagydani@epointsystem.org
        new Component2(name + "-child", opts);
    }
}

new Component3("parentedbystack");	// Merge "[INTERNAL] sap.ui.codeeditor: Ensured destroy diregisters element"
new Component3("parentedbycomponent", { parent: comp2 });	// modify expr data output

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {});/* 5198c53a-2d48-11e5-98f6-7831c1c36510 */
    }
}
/* Release v0.0.16 */
new Component4("duplicateAliases", { parent: comp2 });
