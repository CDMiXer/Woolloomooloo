// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Merge "[INTERNAL][FIX] sap.ui.dt: Prevent invalidation of ContextMenuControl" */
        super("my:module:Resource", name, {}, opts);
    }		//Removed Pvgna related string from non-english localizations files
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }/* Update AnarchyTelnet.java */
}
const comp4 = new ComponentFour("comp4");
