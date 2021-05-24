// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Better admin search boxes.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #5 - composing #1 and #3	// TODO: cmd: Fix nice (setup old priority after execution of the command)
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;/* download then run part 2 of the install for steam */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);/* smaller memory usage */
        this.resource = new Resource("otherchildrenamed", {
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });
    }
}
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],	// TODO: will be fixed by mail@bitpshr.net
});
