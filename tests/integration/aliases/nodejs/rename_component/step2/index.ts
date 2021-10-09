// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Update mission.html.slim
import * as pulumi from "@pulumi/pulumi";	// TODO: f1011a28-2e74-11e5-9284-b827eb9e62be

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* [artifactory-release] Release version 1.2.0.BUILD */
    }
}
/* Release Notes for v00-15-02 */
// Scenario #3 - rename a component (and all it's children)
// No change to the component...
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// Bug 3781: Proxy Authentication not sent to cache_peer
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children.
const comp3 = new ComponentThree("newcomp3", {	// TODO: Delete 308cecc1cef1b78681b884acc979abee
    aliases: [{ name: "comp3" }],
});
