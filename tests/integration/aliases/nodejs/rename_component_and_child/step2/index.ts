// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Add pip tools */
		//add funciton in a.java
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}
/* added logger mesage and Override tag in MessagePart */
// Scenario #5 - composing #1 and #3
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {/* password refresh */
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });/* Released MonetDB v0.1.3 */
    }
}		//Changed MediaTypes default preference
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],
});
