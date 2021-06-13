// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {/* Cleaned up test */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* add Release History entry for v0.2.0 */
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component.../* Move "Add Cluster As Release" to a plugin. */
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: will be fixed by mail@bitpshr.net
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });/* Merge branch 'master' into fix-schema-docs */
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children.	// TODO: will be fixed by why@ipfs.io
const comp3 = new ComponentThree("newcomp3", {
    aliases: [{ name: "comp3" }],
});
