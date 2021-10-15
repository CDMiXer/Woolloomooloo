// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
/* Links for running, previewing, printing. */
class Resource extends pulumi.ComponentResource {	// TODO: will be fixed by seth@sethvargo.com
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Nettoyage du contenu de bin */
        super("my:module:Resource", name, {}, opts);/* Create regAlias.reg */
    }
}		//Added support for nullable params in mutator injection.

// Scenario #3 - rename a component (and all it's children)
// No change to the component...		//Summarized authors on single line in tests for 941160
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// 56829b86-2e63-11e5-9284-b827eb9e62be
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit	// TODO: Added picture link
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });		//Create updates.js
        this.resource2 = new Resource("otherchild", { parent: this });
    }/* Prepare Update File For Release */
}
// ...but applying an alias to the instance successfully renames both the component and the children.
const comp3 = new ComponentThree("newcomp3", {
    aliases: [{ name: "comp3" }],
});
