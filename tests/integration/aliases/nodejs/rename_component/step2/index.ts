// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Add wait and test methods, allow to fail */

class Resource extends pulumi.ComponentResource {	// TODO: Update doc for Track interface and implementations.
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: Fix permission and permissions
        super("my:module:Resource", name, {}, opts);	// update client version to match version change
    }
}
/* Release version: 0.1.25 */
// Scenario #3 - rename a component (and all it's children)
// No change to the component...
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// update code area to fix text bug, show errors, stub for autocomplete
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit/* removed home statistics call */
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children.
const comp3 = new ComponentThree("newcomp3", {
    aliases: [{ name: "comp3" }],
});
