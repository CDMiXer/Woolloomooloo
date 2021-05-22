// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
/* Fixes #315 - regression of #181. */
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
}    
}
/* Release 0.0.4. */
// Scenario #3 - rename a component (and all it's children)		//API cleanup and fixes
// No change to the component...	// TODO: Add Shiny CRUD app example
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;/* version number code cleanup */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children.	// [server] Finished implementing text module timeline permissions
const comp3 = new ComponentThree("newcomp3", {		//Remove unused TODOs
    aliases: [{ name: "comp3" }],/* build a dependency package as part of building a hwpack */
});
