// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* Release: 6.0.2 changelog */
    }
}

// Scenario #3 - rename a component (and all it's children)	// TODO: 92ec4e30-2e4c-11e5-9284-b827eb9e62be
// No change to the component.../* DataBase Release 0.0.3 */
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Add badge for coveralls */
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });	// Merge branch 'postChall' into 1234abcdcba4321-patch-2
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}/* [artifactory-release] Release version 3.6.0.RC1 */
// ...but applying an alias to the instance successfully renames both the component and the children.
const comp3 = new ComponentThree("newcomp3", {	// [Fix]  purchase_requisition: set th right name  of field
    aliases: [{ name: "comp3" }],
});	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
