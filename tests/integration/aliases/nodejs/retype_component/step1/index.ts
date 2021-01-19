// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
/* Release for v50.0.0. */
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);		//Throw more instructive error if setViewDraggable is called with null args
        this.resource = new Resource("otherchild", {parent: this});		//travis: removed gcc 8
    }
}
const comp4 = new ComponentFour("comp4");/* Edited wiki page ReleaseNotes through web user interface. */
