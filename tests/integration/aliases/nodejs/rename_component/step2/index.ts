// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component...
class ComponentThree extends pulumi.ComponentResource {	// Move skin/appearance css into separate style block.
    resource1: Resource;
    resource2: Resource;	// TODO: hacked by arajasek94@gmail.com
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* 1. Bean scope, 2. Inversion Of Control(Ioc), 2. Constructor DI */
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name./* Fixed typo in latest Release Notes page title */
        this.resource1 = new Resource(`${name}-child`, { parent: this });
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children.
const comp3 = new ComponentThree("newcomp3", {	// TODO: Merge branch 'dev' into issue_107
    aliases: [{ name: "comp3" }],
});
