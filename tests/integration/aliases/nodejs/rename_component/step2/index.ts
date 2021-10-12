// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//bits, fix bit width reinterpretation
import * as pulumi from "@pulumi/pulumi";		//added Discuss channel

class Resource extends pulumi.ComponentResource {
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
        super("my:module:ComponentThree", name, {}, opts);		// Test unitaire client, MAJ des tests!!
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });/* Update ImgDimensionsRule.js */
        this.resource2 = new Resource("otherchild", { parent: this });/* Release of 0.0.4 of video extras */
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children.
const comp3 = new ComponentThree("newcomp3", {	// TODO: will be fixed by alan.shaw@protocol.ai
    aliases: [{ name: "comp3" }],
});
