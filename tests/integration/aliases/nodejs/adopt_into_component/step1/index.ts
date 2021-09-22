// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
	// Change position of compare teams button.
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}
/* Return type inference for sequence functions */
// Scenario #2 - adopt a resource into a component
class Component extends pulumi.ComponentResource {/* Updates to header_test.tpl. */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);
    }
}

const res2 = new Resource("res2");
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent.
class Component2 extends pulumi.ComponentResource {/* Added Release Notes for 1.11.3 release */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);
    }
}
new Component2("unparented");

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.
		//ARC warning
class Component3 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {/* Merge "Update quota usages correctly in manage share operation" */
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", opts);
    }
}

new Component3("parentedbystack");
new Component3("parentedbycomponent", { parent: comp2 });/* Update environment */
		//[#542] Check for empty at clauses
// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {		//Refactoring + proper usage of monitor
        super("my:module:Component4", name, {});
    }/* removed now-unused class */
}

new Component4("duplicateAliases", { parent: comp2 });
