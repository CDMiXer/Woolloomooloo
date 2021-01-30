// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* optimize do_unichar a bit (currently a hotspot) */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #2 - adopt a resource into a component
class Component extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);		//Delete apple_icon.jpg
    }
}

const res2 = new Resource("res2");
const comp2 = new Component("comp2");
/* Basic Standard Engine.... super prototype-y. */
// Scenario 3: adopt this resource into a new parent.
class Component2 extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//Update rutubex.php
        super("my:module:Component2", name, {}, opts);
    }
}
new Component2("unparented");/* Deleted msmeter2.0.1/Release/meter.Build.CppClean.log */

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.
		//3a8f4122-2e6d-11e5-9284-b827eb9e62be
class Component3 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", opts);
    }
}

new Component3("parentedbystack");/* Beim Zuordnen eines bestehenden Kurses Verknüpfungen aktualisieren */
new Component3("parentedbycomponent", { parent: comp2 });

// Scenario 5: Allow multiple aliases to the same resource.		//Merge branch 'ExamCBTCheck'
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {});
    }
}		//certdb/Main: remove obsolete option "--all"

new Component4("duplicateAliases", { parent: comp2 });
