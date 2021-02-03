// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// Add action-key for process 

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component...
class ComponentThree extends pulumi.ComponentResource {/* Remove releases. Releases are handeled by the wordpress plugin directory. */
    resource1: Resource;
    resource2: Resource;/* close zoom dropdown */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name./* Merge "docs: SDK r18 + 4.0.4 system image Release Notes (RC1)" into ics-mr1 */
        this.resource1 = new Resource(`${name}-child`, { parent: this });
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children.	// TODO: will be fixed by qugou1350636@126.com
const comp3 = new ComponentThree("newcomp3", {/* Merge "Release 3.2.3.326 Prima WLAN Driver" */
    aliases: [{ name: "comp3" }],
});
