// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }/* Final Release Creation 1.0 STABLE */
}	// TODO: Fixes the most annoying thing about admin helping.

// Scenario #2 - adopt a resource into a component
class Component extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);
    }
}

const res2 = new Resource("res2");
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent.
class Component2 extends pulumi.ComponentResource {	// rev 532867
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);
    }	// TODO: Adjust typo license file.
}
new Component2("unparented");

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix	// TODO: will be fixed by steven@stebalien.com
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", opts);/* Updated Release Notes for the upcoming 0.9.10 release */
    }/* Delete test_bug.cc */
}

new Component3("parentedbystack");
new Component3("parentedbycomponent", { parent: comp2 });

// Scenario 5: Allow multiple aliases to the same resource.	// TODO: will be fixed by hugomrdias@gmail.com
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {});	// TODO: hacked by brosner@gmail.com
    }
}

new Component4("duplicateAliases", { parent: comp2 });
