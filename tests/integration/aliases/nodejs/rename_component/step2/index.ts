// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Release 1.5.0 */

class Resource extends pulumi.ComponentResource {/* Merge "Add affiliations for ijsnellf" */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component...
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children./* VisDialog style fix. */
const comp3 = new ComponentThree("newcomp3", {/* Merge "Release notes - aodh gnocchi threshold alarm" */
    aliases: [{ name: "comp3" }],
});	// TODO: hacked by davidad@alum.mit.edu
