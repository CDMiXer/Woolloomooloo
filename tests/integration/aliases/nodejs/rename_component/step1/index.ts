// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release notes for 2.8. */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {		//change addons
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}/* Removed ormlite refs. */

// Scenario #3 - rename a component (and all it's children)/* FRESH-329: Update ReleaseNotes.md */
class ComponentThree extends pulumi.ComponentResource {/* Add AF to graphics settings */
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, {parent: this});
        this.resource2 = new Resource("otherchild", {parent: this});
    }/* Release of eeacms/bise-backend:v10.0.32 */
}
const comp3 = new ComponentThree("comp3");
