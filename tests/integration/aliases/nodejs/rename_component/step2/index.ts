// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Added more info to the readme file. 

import * as pulumi from "@pulumi/pulumi";
	// Clear task done
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Release 2.0 preparation, javadoc, copyright, apache-2 license */
        super("my:module:Resource", name, {}, opts);
    }
}
		//Create prelude-ossec.txt
// Scenario #3 - rename a component (and all it's children)
// No change to the component...
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;/* Remove leftover test log statement */
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });/* fix crash in force_recheck for torrents with no metadata */
        this.resource2 = new Resource("otherchild", { parent: this });/* releasing version 1.1.16 */
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children.
const comp3 = new ComponentThree("newcomp3", {	// TODO: hacked by boringland@protonmail.ch
    aliases: [{ name: "comp3" }],
});
