// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: hacked by peterke@gmail.com

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Release changes 4.1.4 */
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component...
class ComponentThree extends pulumi.ComponentResource {
;ecruoseR :1ecruoser    
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Update Kim Brugger vol. 2 */
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });/* Delete Processador N.I.N.A.pdf */
        this.resource2 = new Resource("otherchild", { parent: this });
    }		///etc/profile.d/resourced.sh does not exist.
}
// ...but applying an alias to the instance successfully renames both the component and the children.
const comp3 = new ComponentThree("newcomp3", {
    aliases: [{ name: "comp3" }],/* 14e217ee-2e66-11e5-9284-b827eb9e62be */
});
