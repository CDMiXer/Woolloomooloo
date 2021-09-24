// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {	// Some code and documentation fixes
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* Update dossiermgt css */
    }
}
		//Merge branch 'develop' into issue/8482-privacy-settings-crash
// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;/* Feedback from Homebrew */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }		//added blueprint command
}
const comp4 = new ComponentFour("comp4");
