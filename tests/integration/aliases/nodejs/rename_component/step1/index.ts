// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }	// feint, tailwind, metal burst (new moves)
}

// Scenario #3 - rename a component (and all it's children)/* [artifactory-release] Release version 3.1.13.RELEASE */
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;		//Update version and changelog 3.1
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.		//Fix bug in ParametricPlot y-range; F.show() show multiple JSFormData's
        this.resource1 = new Resource(`${name}-child`, {parent: this});
        this.resource2 = new Resource("otherchild", {parent: this});
    }
}
const comp3 = new ComponentThree("comp3");/* Update Data_Portal_Release_Notes.md */
