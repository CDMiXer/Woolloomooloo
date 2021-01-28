// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #5 - composing #1 and #3		//some readme tweaks
class ComponentFive extends pulumi.ComponentResource {	// Move main source folder
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {/* Document the Job controller. */
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],	// TODO: Temporarily use Flutter master branch instead of dev
        });
    }
}	// TODO: will be fixed by davidad@alum.mit.edu
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],
});
