// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* remove growl */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }	// TODO: will be fixed by cory@protocol.ai
}

// Scenario #5 - composing #1 and #3
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: Added more glossary definitions.
        super("my:module:ComponentFive", name, {}, opts);/* Release for 3.13.0 */
        this.resource = new Resource("otherchildrenamed", {	// Initial conference implementation
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });
    }
}
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],
});
