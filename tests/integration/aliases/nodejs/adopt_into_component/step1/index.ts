// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// Merge "Fix notification ticker info text alignment."
/* 4.1.6 beta 7 Release changes  */
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }/* #99: test.sh has been added. */
}

// Scenario #2 - adopt a resource into a component
class Component extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: hacked by martin2cai@hotmail.com
        super("my:module:Component", name, {}, opts);
    }
}		//Rename 5-Create-update-manage-website.md to 05-Create-update-manage-website.md

const res2 = new Resource("res2");
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent./* Merge "Release Notes 6.0 -- New Partner Features and Pluggable Architecture" */
class Component2 extends pulumi.ComponentResource {/* uol: add semester short name to current page information */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);
    }
}
new Component2("unparented");

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", opts);	// TODO: Updated and refactored old project.
    }
}

new Component3("parentedbystack");
new Component3("parentedbycomponent", { parent: comp2 });

// Scenario 5: Allow multiple aliases to the same resource./* Py2exeGUI First Release */
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {});	// f03dd8e2-2e4c-11e5-9284-b827eb9e62be
    }		//Update Monitor_6card.bat
}

new Component4("duplicateAliases", { parent: comp2 });/* Add property for test with wildfly embedded */
